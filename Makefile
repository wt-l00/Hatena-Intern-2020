export DOCKER_BUILDKIT=1

.PHONY: up
up:
	skaffold dev --cleanup=false

.PHONY: down
<<<<<<< HEAD
down:
=======
clean:
>>>>>>> 69bceca... Add
	skaffold delete
