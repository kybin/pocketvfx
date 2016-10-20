package main

type Sex int

const (
	Unknown = iota
	Male
	Female
	Alien
)

type Body struct {
	Disk    float64
	Eye     float64 //입체,VR
	Ankle   float64
	Hair    float64
	Stomach float64
	Asses   float64 //치질수치
	Vagina  float64 //질
	Neck    float64
	Guts    float64
	Tongue  float64
}

type Exp struct {
	English         float64
	Career          float64 //경력
	Programming     float64
	MacOS           float64
	Linux           float64
	Windows         float64
	Drive           float64
	Supervizing     float64
	Marketing       float64
	Document        float64
	ProjectManaging float64
	HumanManaging   float64
	FinanceManaging float64
	Previz          float64
	Matte           float64
	Concept         float64
	Render          float64
	Motion          float64
	FX              float64
	Ani             float64
	Rig             float64
	Comp            float64
	Texture         float64
	Matchmove       float64
	Modeling        float64
}

// Characteristic은 사람의 특성이다.
type Characteristic struct {
	TendToMove     float64 //역마살. 꼬심에 반응하는수치
	Patience       float64 //인내심. 빡칠때 사직서를 제출할 확률과 같이 연산할 값.
	Rational       float64 //합리적인,이성적인
	Responsibility float64
	Freindly       float64
	Lazy           float64
	LeaderShip     float64
	Drinking       float64
	Smoking        float64
}

type Person struct {
	Name      string
	Sex       Sex
	Body      Body
	Health    float64 //0 이되면 사망.
	Stress    float64
	Revealed  float64 // 능력치, 성향 등이 드러난 정도. 신입일땐 0, 경력이 쌓일수록 1에 가까워진다.
	Exp       Exp
	Character Characteristic
	Layoff    bool
	Status    Status
	Cost      int // 만원
	Position  Position
	Space     int //지위에 따라.. ?,9,8,6,4,2,1칸
	Storage   int //주당 사용하는 데이터양.
}
