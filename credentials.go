// credentials
package main

type Credentials struct {
	PrivateKeyId string
	PrivateKey   string
	ClientId     string
	Type         string
	ISS          string
}

var credentials = Credentials{
	PrivateKeyId: "b69bc0d8d547aeeed972f6e24b9f285d24972780",
	PrivateKey:   "-----BEGIN PRIVATE KEY-----\nMIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAOGIlzb0xkeyrEbo\nJKGK6Xu3N+D4ErQPDwZHwc2cu+aJztBKTG/d9wQs8lM24tmdUfC080fD+bJ4Ll1I\nS7mu4yM0yAciYHzlSMhMGoWGPoYVIV+RMD9zV1MFk6/QkFCzPQiofF1j27qGR+sP\ngz7SrT6SW2kbGpVROW7NxGpA5psBAgMBAAECgYEA1T29o+OJ+k6QnzZtNCRQoH3a\noqZLoN3AGfMM5ii0bMSNu9kULo1TOip5MH7rvFAIIqwkfg1o8Pq6884gcz9NV7zS\nCjCfU7GwfYxXlwNJP0ro2dyVMoOejil7p8ISgY05rE1neeUsQMRYCd4MtkB6MFFa\nYmt8umIEQYA2YndpizUCQQD2PxemAxPf1TQC20tE3WmaDBt3/qpF4Og5YD4DaiAH\nzr+C54uIgE2OM5O37dWgyBktrHsCdwgaau4/sQXvB8evAkEA6nd44C9NiyREuZlP\nfKrf1wjFAv/viy6qH4CeD9Hut7h+yLNuksflFxPwh1PR9ABV3lUj0i+kgO7eiIgs\nqsLETwJBAIBXH88xeXFGoIYAn3kz4O8d8k8XFs22y7HjvE5xSGJpWh0y6uFo5YMA\n1LOfLKFQyPaqF66QNVP3eVcOv2X2axcCQQDKTrP5nFcoS+8QxfwT8cyaWnLDl9o1\nGdqou+2mcNmtpH+g5VHvTJObShUsb2KlSvTMlmIGJh2nMnTJEdaYsN15AkAwEoeD\noCZpOvT1yhrcClO+6O1kpVRtuQVX/yWT6MKuDh5mdI0lhcOuLleMDsQC4niLvRRH\nxQrd48otOlvxfmZL\n-----END PRIVATE KEY-----\n",
	ClientId:     "779906410321-j2lbks58oumq4jjlil1igqijqj6a9ofu.apps.googleusercontent.com",
	Type:         "service_account",
	ISS:          "779906410321-j2lbks58oumq4jjlil1igqijqj6a9ofu@developer.gserviceaccount.com",
}
