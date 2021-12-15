package particles

import (
	"project-particles/config"
	"math/rand"
	"log"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {

	var particules []Particle
	var s System = System{}

	if config.General.RandomSpawn {
		s.getPos = getRandomPosition
	}else {
		s.getPos = getFixedPosition
	}
		
	for i  := 0; i < config.General.InitNumParticles; i++ {
		var x, y = s.getPos()
		var p Particle = genParticule(x, y)
		particules = append(particules, p)
	}

	
	s.UpdateCount = 0
	s.Content = particules

	if config.General.Debug {
		log.Println("Nouveau système créé avec pour contenu :", s.Content)
	}

	return s
}

func (s *System) UpdateContent() {
	if config.General.SpawnRate > 0 {
		if config.General.SpawnRate > 1 {
			for i := 0; i < int(config.General.SpawnRate); i++ {
				var x, y = s.getPos()
				s.Content = append(s.Content, genParticule(x, y))
			}
		} else {
			s.UpdateCount += config.General.SpawnRate
			if s.UpdateCount >= 1 {
				s.UpdateCount = 0
				var x, y = s.getPos()
				s.Content = append(s.Content, genParticule(x, y))
			}
		}
		
	}
}

func getRandomPosition() (float64, float64) {
	return float64(rand.Intn(config.General.WindowSizeX)), float64(rand.Intn(config.General.WindowSizeY))
}

func getFixedPosition() (float64, float64) {
	return float64(config.General.SpawnX), float64(config.General.SpawnY)
}

func genParticule(x, y float64) Particle {
	return Particle{
		PositionX: x, 
		PositionY: y,
		VitesseX: rand.Float64() * 11 - 5,
		VitesseY: rand.Float64() * 11 - 5,
		ScaleX: 1,
		ScaleY: 1,
		ColorRed: 1,
		ColorGreen: 0,
		ColorBlue: 0,
		Opacity: 1,
	}
}