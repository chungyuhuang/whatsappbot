package constant

import "os"

var (
	Receiver = os.Getenv("RECEIVER_ID")
	// "6592413940-1535172864@g.us"

	PayRentalMsg = "各位 繳房租囉！ 感謝"
	PayInternetFeeMsg = "各位 繳網路費10塊囉！ 感謝"
)
