.PHONY : install_and_run install run clean

full_run: install run

install:
	$(MAKE) clean
	docker build -t a-tour-of-go .

run:
	./docker/restart-docker-image.sh a-tour-of-go

clean:
	./docker/clean-docker-images.sh
