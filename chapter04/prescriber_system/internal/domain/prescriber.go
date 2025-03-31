package domain

import (
	"context"
	"errors"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

var (
	ErrQueueFull          = errors.New("處方隊列已滿，請稍後再試")
	ErrRequestTimeout     = errors.New("處方請求超時")
	ErrCreatePrescription = errors.New("創建處方失敗")
)

// PrescriptionRequest represents a request for prescription
type PrescriptionRequest struct {
	Ctx    context.Context
	Demand PrescriptionDemand
	Result chan Prescription
	Error  chan error
}

// Prescriber represents a medical professional who can diagnose and prescribe
type Prescriber struct {
	registry          PrescribeHandlerRegistry
	patientDatabase   *PatientDatabase
	prescribeHandler  PrescribeHandler
	prescriptionQueue chan PrescriptionRequest
	subscribers       []Subscriber
	mu                sync.Mutex
	wg                sync.WaitGroup
	ctx               context.Context
	cancel            context.CancelFunc
}

// NewPrescriber creates a new prescriber with the given information
func NewPrescriber(
	ctx context.Context,
	registry PrescribeHandlerRegistry,
	patientDatabase *PatientDatabase,
) *Prescriber {
	ctx, cancel := context.WithCancel(ctx)

	p := &Prescriber{
		registry:          registry,
		patientDatabase:   patientDatabase,
		prescriptionQueue: make(chan PrescriptionRequest, 100),
		subscribers:       make([]Subscriber, 0),
		ctx:               ctx,
		cancel:            cancel,
	}

	// Start the prescription processor
	go p.processPrescriptionQueue()

	return p
}

func (p *Prescriber) LoadPrescribeList(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	prescribes := strings.Split(strings.TrimSpace(string(data)), "\n")
	p.CreatePrescribeChain(prescribes)

	return nil
}

func (p *Prescriber) CreatePrescribeChain(prescribes []string) {
	var head PrescribeHandler
	var prev PrescribeHandler

	for _, prescribe := range prescribes {
		prescribeHandler := p.registry.Find(prescribe)
		if prescribeHandler == nil {
			continue
		}

		// 串接 Handler 責任鏈
		if head == nil {
			head = prescribeHandler
		} else {
			prev.SetNext(prescribeHandler)
		}
		prev = prescribeHandler
	}

	p.prescribeHandler = head
}

// Prescribe adds a prescription request to the queue with context
func (p *Prescriber) Prescribe(ctx context.Context, demand PrescriptionDemand) (Prescription, error) {
	resultChan := make(chan Prescription, 1)
	errorChan := make(chan error, 1)

	request := PrescriptionRequest{
		Ctx:    ctx,
		Demand: demand,
		Result: resultChan,
		Error:  errorChan,
	}

	if err := p.enqueue(request); err != nil {
		return Prescription{}, err
	}

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errorChan:
		return Prescription{}, err
	case <-ctx.Done():
		return Prescription{}, ErrRequestTimeout
	}
}

// enqueue safely adds request to queue
func (p *Prescriber) enqueue(request PrescriptionRequest) error {
	select {
	case p.prescriptionQueue <- request:
		return nil
	default:
		return ErrQueueFull
	}
}

// processPrescriptionQueue continuously processes prescription requests
func (p *Prescriber) processPrescriptionQueue() {
	for {
		select {
		case request, ok := <-p.prescriptionQueue:
			if !ok {
				return
			}

			// 追蹤 Goroutine
			p.wg.Add(1)
			defer p.wg.Done()

			p.handlePrescription(request)

		case <-p.ctx.Done():
			return
		}
	}
}

// handlePrescription processes a single prescription request
func (p *Prescriber) handlePrescription(req PrescriptionRequest) {
	start := time.Now()

	p.mu.Lock()
	defer p.mu.Unlock()

	// Retrieve patient from database
	patient, err := p.patientDatabase.FindByID(req.Demand.PatientID)
	if err != nil {
		req.Error <- err
		return
	}

	// Handle prescription
	prescription, err := p.prescribeHandler.HandlePrescribe(*patient, req.Demand.GetSymptoms())
	if err != nil {
		req.Error <- errors.Join(ErrCreatePrescription, err)
		return
	}

	// 確保處理時間至少 3 秒
	elapsed := time.Since(start)
	if elapsed < 3*time.Second {
		time.Sleep(3*time.Second - elapsed)
	}

	req.Result <- prescription
}

// Shutdown gracefully stops the prescriber
func (p *Prescriber) Shutdown() {
	// 停止接受新請求
	p.cancel()

	// 等待所有請求完成
	p.wg.Wait()

	// 關閉請求隊列
	close(p.prescriptionQueue)
}

// 訂閱管理
func (p *Prescriber) AddSubscriber(sub Subscriber) {
	p.subscribers = append(p.subscribers, sub)
}

func (p *Prescriber) RemoveSubscriber(sub Subscriber) {
	for i, s := range p.subscribers {
		if s.Name() == sub.Name() {
			p.subscribers = slices.Delete(p.subscribers, i, i+1)
			break
		}
	}
}
