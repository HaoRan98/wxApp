package util

import (
	"strings"
)

func SWJG(deptid string) (string, string , string) {

	if deptid == "1370600" {
		return "13706000000", "烟台市税务局","市局"
	} else if deptid == "1370692" {
		return "13706920000", "烟台保税港区税务局","保税港区"
	} else if deptid == "1370686" {
		return "13706860000", "栖霞市税务局","栖霞市"
	} else if deptid == "1370682" {
		return "13706820000", "莱阳市税务局","莱阳市"
	} else if deptid == "1370684" {
		return "13706840000", "蓬莱市税务局","蓬莱市"
	} else if deptid == "1370685" {
		return "13706850000", "招远市税务局","招远市"
	} else if deptid == "1370683" {
		return "13706830000", "莱州市税务局","莱州市"
	} else if deptid == "1370693" {
		return "13706930000", "烟台高新技术产业开发区税务局","高新区"
	} else if deptid == "1370691" {
		return "13706910000", "烟台经济开发区税务局","开发区"
	} else if deptid == "1370634" {
		return "13706340000", "烟台市长岛县税务局","长岛县"
	} else if deptid == "1370612" {
		return "13706120000", "烟台市牟平区税务局","牟平区"
	} else if deptid == "1370681" {
		return "13706810000", "龙口市税务局","龙口市"
	} else if deptid == "1370611" {
		return "13706110000", "烟台市福山区税务局","福山区"
	} else if deptid == "1370687" {
		return "13706870000", "海阳市税务局","海阳市"
	} else if deptid == "1370613" {
		return "13706130000", "烟台市莱山区税务局","莱山区"
	} else if deptid == "1370602" {
		return "13706020000", "烟台市芝罘区税务局","芝罘区"
	} else {
		return "", "",""
	}

}

func QuDiaoSWJ(deptname string) string {
	deptname = strings.Replace(deptname,"税务局","",-1)
	return deptname
}
