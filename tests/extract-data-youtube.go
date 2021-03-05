package tests

import (
	"fmt"
	djyoutube "github.com/pkanti/v2/youtube"
)

func TestGetMetadataYoutube() {
	fmt.Println("Extracting information from https://www.youtube.com/playlist?list=OLAK5uy_lgDycc9EqG0S7P_M0SHmwl3UuYWQ4JRqg")
	fmt.Println(djyoutube.GetInfo("https://www.youtube.com/playlist?list=OLAK5uy_lgDycc9EqG0S7P_M0SHmwl3UuYWQ4JRqg"))
	fmt.Println("Extracting information from https://music.youtube.com/watch?v=stj5Wijk-0s&feature=share")
	fmt.Println(djyoutube.GetInfo("https://music.youtube.com/watch?v=stj5Wijk-0s&feature=share"))
	fmt.Println("Extracting information from https://music.youtube.com/playlist?list=OLAK5uy_mTefpmawCkdsfMPIlAXQbSzNFeak0MVOI")
	fmt.Println(djyoutube.GetInfo("https://music.youtube.com/playlist?list=OLAK5uy_mTefpmawCkdsfMPIlAXQbSzNFeak0MVOI"))
	fmt.Println("Extracting information from youtu.be/CznH00UBwQU")
	fmt.Println(djyoutube.GetInfo("youtu.be/CznH00UBwQU"))
}
