// main.go
package main

import (
	"encoding/json"
	"log"
)

func main() {
	bi1 := Service{"COM设计", "调阅历史工作文件、账目，审议各部门工作计划、工作报告，以电话、会议、面谈等方式了解需求，以书面方式提交设计报告。可以根据审议结果修改一个版本。不超过四天：人民币贰万元，第五天起每天九千六百元。注：提前一个季度预约。工作地点北京，如需离京另收差旅费，数额协商决定。", 20000}
	bi2 := Service{"COD部署", "根据P2俱乐部成员设计的COM方案，考虑团队成员意见、现有设备设施情况以及当地政府部门的工作尺度，以书面方式提交具体部署方案。对律师工作结果提出书面审议意见。不超过十天人民币六万元，第十一天起每天一万元。注：提前一个季度预约。工作地点北京，如需离京另收差旅费，数额协商决定。", 60000}
	huangyg := MemberInfo{"huangyg", "管理员", "高级会员", false, "", "", "1,2,5,6,7,9,11,12", "huangyg@p2club.org", 0, 903, []Service{bi1, bi2}}
	huangygInfo, _ := json.Marshal(huangyg)

	str := string(huangygInfo[:])
	log.Print(str)

	ia := InformationAssets{"会员信息", "huangyg", str, "-----BEGIN PGP SIGNATURE-----Version: PGP Desktop 10.2.0 (Build 2599) - not licensed for commercial use: www.pgp.comCharset: utf-8wsBVAwUBVHS4XPb4AsDNppvZAQhq1QgAksGBqcBK00ahNchKdIKK/NVUdUqYGlgf3YZEnIUWjGLECet/E9UoiT0/u57NzbOiy+klN0u1Q7yvuuLlDPjRsAscpVAX80r7imM33oCbQy2AGIhA0QE3pFS4sYAgLsOWsjrO4HZ5glKIyN4O06zcDH4Qn2BMJYBRHN5zwYOggrS3xUQdt0agQpgUmDisCIjGZyqrM/8/HcBlIVLpPO+ICHSsAz3WY9Dj7eTZG3HvhHLFIwVoyuYhQyu0W0tjVcnPiSN7d5tlUOPq6LUgGFQmZmu9X45QhmNtK2uefsCpyCpXpmhD3aYzUPL3UqDui1Oj0YNJ9SJDJy8dMD7VVbR7zg===y5B7-----END PGP SIGNATURE-----"}
	RetJson, _ := json.Marshal(ia)

	log.Printf("%s", RetJson)
}
