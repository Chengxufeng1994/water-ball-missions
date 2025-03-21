package main

func main() {
	// 創建頻道
	pewdiepie := NewChannel("PewDiePie")
	waterBallAcademy := NewChannel("WaterBallAcademy")

	//創建訂閱者
	waterBall := NewWaterBallSubscriber()
	fireball := NewFireBallSubscriber()

	// 訂閱頻道
	pewdiepie.AddSubscriber(waterBall)
	pewdiepie.AddSubscriber(fireball)
	waterBallAcademy.AddSubscriber(waterBall)
	waterBallAcademy.AddSubscriber(fireball)

	pewdiepie.Upload(NewVideo("Hello guys", "Clickbait", 30))
	waterBallAcademy.Upload(NewVideo("C1M1S3、敘述：物件 vs. 類別", "Clickbait", 600))
	pewdiepie.Upload(NewVideo("Minecraft", "Let's play Minecraft", 1800))
}
