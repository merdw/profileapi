package api

import (
	"fmt"
	"github.com/Davincible/goinsta/v3"
)

var insta *goinsta.Instagram
var profil *goinsta.Profile
var follower, following int
var bio, name string
var pplink string
var isverified bool
var eror bool = false

func Jsonimport() *goinsta.Instagram {
	insta, _ = goinsta.Import("./test.json")
	defer insta.Export("test.json")
	return insta
}

func VisitProfile(nickname string) *goinsta.Profile {
	Jsonimport()

	profil, _ = insta.VisitProfile(nickname)

	return profil
}

func NameProfile(nickname string) (int, int, string, string, string, bool, bool) {
	Jsonimport()
	test, err := insta.Profiles.ByName(nickname)

	if err != nil {

		fmt.Println("HATA:", err)
		eror = true
		follower = 0
		following = 0
		bio = "bos"
		name = "bos"

		pplink = "bos"

		isverified = false

		return follower, following, bio, name, pplink, isverified, eror
	} else {
		usr := test

		follower = usr.FollowerCount
		following = usr.FollowingCount
		bio = usr.Biography
		name = usr.FullName
		pplinkjson := usr.HdProfilePicURLInfo
		pplink = pplinkjson.URL

		isverified = usr.IsVerified

		return follower, following, bio, name, pplink, isverified, eror
	}

}

func GetInfo(nickname string) (int, int, string, string, string, bool) {
	VisitProfile(nickname)

	usr := profil.User
	follower = usr.FollowerCount
	following = usr.FollowingCount
	bio = usr.Biography
	name = usr.FullName
	pplinkjson := usr.HdProfilePicURLInfo
	pplink = pplinkjson.URL

	isverified = usr.IsVerified
	return follower, following, bio, name, pplink, isverified
}
