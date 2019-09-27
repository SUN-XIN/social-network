package main

import "log"

var (
	testNormalProfile, testSpecialProfile               *Profile
	testSessionNormalProfile, testSessionSpecialProfile *Session
)

func seedGetProfiles(rout *Router) error {
	testNormalProfile = NewProfile("getprofiletest_name1")
	testNormalProfile.SetStatus(StatusNormal)
	err := rout.DB.InsertProfile(testNormalProfile)
	if err != nil {
		return err
	}

	testSpecialProfile = NewProfile("getprofiletest_name2")
	testSpecialProfile.SetStatus(StatusSpecial)
	err = rout.DB.InsertProfile(testSpecialProfile)
	if err != nil {
		return err
	}

	testSessionNormalProfile = NewSession(testNormalProfile.ID)
	err = rout.DB.InsertSession(testSessionNormalProfile)
	if err != nil {
		return err
	}

	testSessionSpecialProfile = NewSession(testSpecialProfile.ID)
	err = rout.DB.InsertSession(testSessionSpecialProfile)
	if err != nil {
		return err
	}

	log.Printf("---------------------------------------------------- \n")
	log.Printf("---------------------------------------------------- \n")
	log.Printf("---------------------------------------------------- \n")
	log.Printf("test data is inserted -> \n")
	log.Printf("normal profile: %+v \n", testNormalProfile)
	log.Printf("---------------------------------------------------- \n")
	log.Printf("special profile: %+v \n", testSpecialProfile)
	log.Printf("---------------------------------------------------- \n")
	log.Printf("normal profile's session: %+v \n", testSessionNormalProfile)
	log.Printf("---------------------------------------------------- \n")
	log.Printf("special profile's session: %+v \n", testSessionSpecialProfile)
	log.Printf("---------------------------------------------------- \n")
	log.Printf("---------------------------------------------------- \n")
	log.Printf("---------------------------------------------------- \n")
	log.Printf("---------------------------------------------------- \n")

	return nil
}
