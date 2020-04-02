package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	HEIGHT = 15
	WIDTH  = 15
)

//POINT
//think about type Point in import "image"
type Point struct {
	x, y int
}

func getSector(p Point) int {
	zone := math.Ceil(float64(p.x+1)/5.0) + math.Floor(float64(p.y/5.0)*3)
	return int(zone)
}

//DIRECTION
type Direction struct {
	dx   int
	dy   int
	name string
}

func (s *State) getTile(t Tile) (Tile, error) {
	if t.pos.x < 0 || t.pos.x > WIDTH || t.pos.y < 0 || t.pos.y > HEIGHT || t.what != "." {
		return Tile{}, fmt.Errorf("out of bound or island\n") //an error here like out of range or island!!
	} else {
		return s.carte[t.pos.x][t.pos.y], nil
	}
}

//TILE
//getTile() with limit checking
//to see what happen if border crossed?
//https://rosettacode.org/wiki/Bitmap/Flood_fill#Go ==> would like to see the getpx func i imagine a getTile()
// with limit checkers in a Grid Type...
type Tile struct {
	pos  Point
	what string
}

func isWalkable(t Tile) bool {
	return t.what == "."
}
func (s *State) computeNeighbours(t Tile) int {
	var neighbours int
	if t.pos.x-1 >= 0 && s.carte[t.pos.x-1][t.pos.y].what == "." {
		neighbours += 1
	}
	if t.pos.x+1 < WIDTH && s.carte[t.pos.x+1][t.pos.y].what == "." {
		neighbours += 1
	}
	if t.pos.y-1 >= 0 && s.carte[t.pos.x][t.pos.y-1].what == "." {
		neighbours += 1
	}
	if t.pos.y+1 < HEIGHT && s.carte[t.pos.x][t.pos.y+1].what == "." {
		neighbours += 1
	}
	return neighbours
}

//OPP
type Opp struct {
	hitPoints       int
	oppDirection    string
	torpedoPos      []Point
	lastTorpedoZone int
	surfaceZone     int
	sonarZone       int
	dirs            string
}

func (o *Opp) parseOppOrders(orders string) {
	//sanitize orders
	st := strings.ReplaceAll(orders, "|", " ")
	//split it
	ord := strings.Split(st, " ")
	for idx, w := range ord {
		if w == "N" || w == "E" || w == "W" || w == "S" {
			o.oppDirection = w
			//a string of ALL dir followed by my opp
			//must include the silent one!!
			o.dirs = o.dirs + w
		}
		//in move n torpedo i cant see torpedo i only see it when he fires it
		if w == "TORPEDO" {
			x, _ := strconv.Atoi(ord[idx+1])
			y, _ := strconv.Atoi(ord[idx+2])
			o.torpedoPos = append(o.torpedoPos, Point{x, y})
		}
		if w == "SURFACE" {
			e, _ := strconv.Atoi(ord[idx+1])
			o.surfaceZone = e
		}
		if w == "SONAR" {
			e, _ := strconv.Atoi(ord[idx+1])
			o.sonarZone = e
		}
		if w == "SILENCE" {
			//to know the gaps
			o.dirs = o.dirs + "X"
		}
	}
}
func (o *Opp) getLastTorpZone() (int, Point) {
	var zone int
	var lastTPos Point
	if len(o.torpedoPos) > 0 {
		if len(o.torpedoPos) == 1 {
			zone = getSector(o.torpedoPos[0])
			lastTPos = o.torpedoPos[0]
		} else {
			zone = getSector(o.torpedoPos[len(o.torpedoPos)-1])
			lastTPos = o.torpedoPos[len(o.torpedoPos)-1]
		}
	}
	o.lastTorpedoZone = zone
	return zone, lastTPos
}

//ME
type Me struct {
	id           int
	currentPos   Tile
	currentDir   string
	hitPoints    int
	visitedTiles map[Tile]bool
	//dirs
	canGoWest  bool
	canGoEast  bool
	canGoNorth bool
	canGoSouth bool
	//Torpedo
	torpedoCooldown int
	canFireTorpedo  bool
	//Sonar
	sonarCooldown int
	canUseSonar   bool
	sonarResult   string
	//Silence
	silenceCooldown int
	canUseSilence   bool
	//Opp
	oppSurfaceHint string
	//comm to send
	commands      []string
	possibleMoves []Tile
}

//STATE
type State struct {
	board         string
	carte         [HEIGHT][WIDTH]Tile
	directions    []Direction
	walkableTiles []Tile
	me            Me
	opp           Opp
	targets       []Tile
}

/*
// sim///
func estimateTrajectory(s State,numTurns int){
	//simply play the game of trajectory right?
	cpState := s //a copy of the state
	for i := 0, i < numTurns, i++{
if cpState.me.currentDir =="N"{
	//MUST find ja simpler way to do that: getTile!!! and a valid one!! with no shit
	//if y+1 etc
	nextPos = cpState.carte[cpState.me.currentDir.pos.x][cpState.me.currentDir.pos.y-1]
	//if next pos is a deadstreet do not go there!!
	if nextPos.pos.x+1 >0 &&


	}

}
}
*/

//YannTt'as 3 mouvements possible, tu floodfill pour chaque, garde celui qui te laisse le plus de cases dispo après move
//for _,dirs := range s.directions{
func (s *State) checkDirections(t Tile) {
	//les pos devraient être checkées ttes ici puis tu checkes direct dans la tile
	//passée en arg (tu passes x-1) plutot que de checker en static ici...
	//FF je cpds pas retour au neigbour..
	if t.pos.x-1 >= 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !s.me.visitedTiles[s.carte[t.pos.x-1][t.pos.y]] {
		s.me.canGoWest = true
		s.me.possibleMoves = append(s.me.possibleMoves, s.carte[t.pos.x-1][t.pos.y])
	}
	if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !s.me.visitedTiles[s.carte[t.pos.x+1][t.pos.y]] {
		s.me.canGoEast = true
		s.me.possibleMoves = append(s.me.possibleMoves, s.carte[t.pos.x+1][t.pos.y])

	}
	if t.pos.y-1 >= 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !s.me.visitedTiles[s.carte[t.pos.x][t.pos.y-1]] {
		s.me.possibleMoves = append(s.me.possibleMoves, s.carte[t.pos.x][t.pos.y-1])

		s.me.canGoNorth = true
	}
	if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !s.me.visitedTiles[s.carte[t.pos.x][t.pos.y+1]] {
		s.me.possibleMoves = append(s.me.possibleMoves, s.carte[t.pos.x][t.pos.y+1])

		s.me.canGoSouth = true
	}
}
func getBestMove(s State) {
	//it takes state as an arg, clone it to sim?
	//determine which direction is best wrt the num of wtiles left and the opp
	//location (indeed where we think he might be thx to the info from surface, sonar
	//or torpedo)
	//should be there in the end
	//s.checkDirections(m.currentPos)
	//determine what commands we must send so that everything gets charged + choose
	//between torp/sonar/silence?
}

//question is: how to triangulate opp pos from his torpedoes?? BFS??
//WRITE A GUESS FUNC TO GUESS ENEMY POS FROM HINTS!!
//I know...but did i grasp the logic??
func (s *State) woodMoves() {
	//calculate the best value
	//var north, south, east, west float64

	if s.me.canGoSouth {
		s.me.currentDir = "S"
		s.me.move("S")
		//doin it that way is real shit
		n := s.computeNeighbours(s.carte[s.me.currentPos.pos.x][s.me.currentPos.pos.y+1])
		south := 1.0 + float64(n)
		log.Println(south, n)
	}
	//what if i can go west?? if ffdW > a ffdE ne devrais je pas ponderer l'éval de W+0.5
	if !s.me.canGoSouth && s.me.canGoEast {
		s.me.currentDir = "E"
		s.me.move("E")
	}
	if !s.me.canGoSouth && !s.me.canGoEast && s.me.canGoNorth {
		s.me.currentDir = "N"
		s.me.move("N")
		/*
			north = 1.0 + ffd["N"]
			if s.me.canGoWest {
				if ffd["W"] > ffd["N"] {
					west += 0.5
				}
			}
		*/
	}
	if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && s.me.canGoWest {
		s.me.currentDir = "W"
		s.me.move("W")
		//west = 1.0 + ffd["W"]
		//c'est une connerie !! je devrais regarder ces ifs c'est une connerie!!
		// je devrais raisonner au niveau de la tile!! combien de direciton possible
		//s il y en a plus qu'une , quelqu'elle soit, c'est que j'ai raté quelque
		//chose!!! et que j'ai plus qu'à perdre une vie en surface!!!
	}
	//log.Println(east, north, south, west)
	//TEST
	//if i am round the torp zone sonar to see what happens
	if getSector(s.me.currentPos.pos) == s.opp.lastTorpedoZone {
		//s.me.sonar(s.opp.lastTorpedoZone) //useless should be able to use sonar in the neighhouring zones
		_, p := s.opp.getLastTorpZone()
		s.me.torpedo(s.carte[p.x][p.y])
	}
	/*
		//Torpedo spam to TEST
		//pb with canFireTorpedo: not enougght charge??
		if !s.me.isTorpCharge() {
			t := s.targets[rand.Intn(len(s.targets))]
			s.me.torpedo(t)
		}
	*/

	/* I surface way too much!! should first learn to use my floodfill to optimize my
	* moves and then i might be able to surface at the right time
	//one direction is possible but cell has already been visited!! so surface
	//must neutralize torpedo here
	if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && !s.me.canGoWest {
		s.me.surface()
		//should reset visited
		for c, _ := range s.me.visitedTiles {
			s.me.visitedTiles[c] = false
		}
	}
	*/
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	//random num generator
	rand.Seed(time.Now().Unix())

	//should have a new func here
	var s State
	s.directions = []Direction{Direction{-1, 0, "W"}, Direction{0, -1, "N"}, Direction{+1, 0, "E"}, Direction{0, +1, "S"}}

	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)
	//did i really care about that??
	s.me.id = myId

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		s.board = s.board + line
	}

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			s.carte[i][j] = Tile{Point{i, j}, string(s.board[j*WIDTH+i])}
			if s.carte[i][j].what == "." {
				s.walkableTiles = append(s.walkableTiles, s.carte[i][j])
			}
		}
	}

	log.Println(s.walkableTiles)
	//my starting pos
	//var startPos = s.walkableTiles[rand.Intn(len(s.walkableTiles))]
	var startPos = s.carte[7][10]
	fmt.Println(startPos.pos.x, startPos.pos.y)

	s.me.visitedTiles = make(map[Tile]bool)

	var turn int
	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
		//should write a New func for State no?
		s.me.currentPos = s.carte[x][y]
		s.me.visitedTiles[s.me.currentPos] = true
		s.me.hitPoints = myLife
		s.me.torpedoCooldown = torpedoCooldown
		s.me.sonarCooldown = sonarCooldown
		s.me.silenceCooldown = silenceCooldown

		s.opp.hitPoints = oppLife

		//TEST TARGET
		dist := s.calculateDist(s.me.currentPos)
		s.getTargets(dist)

		s.checkDirections(s.me.currentPos)
		log.Println(s.me.possibleMoves)
		s.woodMoves()

		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)
		s.me.sonarResult = sonarResult

		scanner.Scan()
		opponentOrders := scanner.Text()
		s.opp.parseOppOrders(opponentOrders)

		s.me.sendTurn()

		//reset turn player data
		//write a reset turn eventually...
		s.me.currentPos = Tile{}
		s.me.canGoNorth = false
		s.me.canGoSouth = false
		s.me.canGoWest = false
		s.me.canGoEast = false
		s.me.possibleMoves = []Tile{}
		s.me.commands = []string{}
		s.targets = []Tile{}
		turn += 1
	}
}
