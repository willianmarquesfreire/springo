package core

type Rights struct {
	Value     int32
	Pos       int8
	Positions []int8
}

var OWNER_READ = Rights{Value: 256, Positions: []int8{8}, Pos: 0}   //100000000
var OWNER_UPDATE = Rights{Value: 128, Positions: []int8{7}, Pos: 1} //010000000
var OWNER_DELETE = Rights{Value: 64, Positions: []int8{6}, Pos: 2}  //001000000

var GROUP_READ = Rights{Value: 32, Positions: []int8{5}, Pos: 3}   //000100000
var GROUP_UPDATE = Rights{Value: 16, Positions: []int8{4}, Pos: 4} //000010000
var GROUP_DELETE = Rights{Value: 8, Positions: []int8{3}, Pos: 5}  //000001000

var OTHER_READ = Rights{Value: 4, Positions: []int8{2}, Pos: 6}   //000000100
var OTHER_UPDATE = Rights{Value: 2, Positions: []int8{1}, Pos: 7} //000000010
var OTHER_DELETE = Rights{Value: 1, Positions: []int8{0}, Pos: 8} //000000001

var DEFAULT_RIGHTS = Rights{
	Value: 504, Positions: []int8{8, 7, 6, 5, 4, 3}, Pos: 6}

var RIGHTS_READ = Rights{
	Value: 504, Positions: []int8{8, 5, 2}, Pos: 6}

var RIGHTS_UPDATE = Rights{
	Value: 504, Positions: []int8{7, 4, 1}, Pos: 6}

var RIGHTS_DELETE = Rights{
	Value: 504, Positions: []int8{6, 3, 0}, Pos: 6}

var OWNER_RIGHTS = Rights{
	Value: 448, Positions: []int8{8, 7, 6}, Pos: 2}



//db.specifications.find({rights: { $bitsAllSet:[8,7,6,5,4,3,2] } }).pretty()  111111100


/**
PESQUISA: db.specifications.find({
			$and:[
			{ui: "wmsystem"},
			{rights: { $bitsAnySet:[8,5,2] } }
			]
		  }).pretty()
 */