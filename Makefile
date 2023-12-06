run:
	@read -p "Enter day number: " day; \
	read -p "Enter part number: " part; \
	echo "Running day$$day/T$$part/main.go"; \
	cd day$$day/T$$part; \
	go run main.go

init-day:
	# copy init/* to dayXX/TX
	@read -p "Enter day number: " day; \
	mkdir -p day$$day; \
	echo "Copying init/* to day$$day"; \
	cp -r init/* day$$day