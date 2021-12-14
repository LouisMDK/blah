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
		var p Particle = Particle{PositionX: x, PositionY: y}
		particules = append(particules, p)
	}

	if config.General.Debug {
		log.Println("Nouveau système créé avec pour contenu :", particules)
	}
	return System{Content: particules}
}
