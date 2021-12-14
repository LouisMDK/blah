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
	
	var getPosition func() (float64, float64)

	// si la variable RandomSpawn dans le fichier config est True, la fonction getPosition() renvoie des valeurs aléatoires
	// sinon elle renvoie les valeurs issues de SpawnX et SpawnY dans le fichier de config
	if config.General.RandomSpawn {
		getPosition = func() (float64, float64) {
			return float64(rand.Intn(config.General.WindowSizeX)), float64(rand.Intn(config.General.WindowSizeY))
		}
	}else {
		getPosition = func() (float64, float64) {
			return float64(config.General.SpawnX), float64(config.General.SpawnY)
		}
	}
		
	for i  := 0; i < config.General.InitNumParticles; i++ {
		var x, y = getPosition()

		var p Particle = genParticule(x, y)
		particules = append(particules, p)
	}

	var f func(*[]Particle)
	if config.General.SpawnRate >= 1 {
		f = func(content *[]Particle) {
			for i := 0; i < int(config.General.SpawnRate); i++ {
				var x, y = getPosition()
				*content = append(*content, genParticule(x, y))
			}
		}
	} else {
		if config.General.SpawnRate > 0 {
			f = func(content *[]Particle) {
				
			}
		}else {
			f = func(content *[]Particle) {

			}
		}
	}
	var s System = System{Content: particules}
	s.UpdateContent = f

	if config.General.Debug {
		log.Println("Nouveau système créé avec pour contenu :", s.Content)
	}

	return s
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