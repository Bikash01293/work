package main
import (
	"fmt"
	"log"
	proto "github.com/golang/protobuf/proto"
)
func main(){
	bikash := &PersonDetails{
		Name : "Bikash",
		Age : 24,
		SocialFollowers : &SocialFollowers{
			Youtube : 1200,
			Twitter : 1400,
		},
	}
	data, err := proto.Marshal(bikash)
	if err != nil{
		log.Fatal("Marshalling error : ", err)
	}
	fmt.Println(data)

	nbikash := &PersonDetails{}
	err = proto.Unmarshal(data, nbikash)
	if err != nil{
		log.Fatal("Unmarshalling error")
	}
	fmt.Println(nbikash.GetName())
	fmt.Println(nbikash.GetAge())
	fmt.Println(nbikash.SocialFollowers.GetYoutube())
	fmt.Println(nbikash.SocialFollowers.GetTwitter())
}

