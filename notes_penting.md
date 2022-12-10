# Gitlab Remote

http://15.1.1.10:2020/dashboard/projects

#Email Ukabima
username : m.aburizal.purnama@ukabima.id
password : BXPSDIEX2O


# Navicate

user : rizal
pass : Navicate1.

# Windows

user : Ideapad
pass : Ukabima69.

# Cockraoch DB
username : rizal
password : 5RqQqKYxQrHJgX0SUkEmrw

# Pass Wifi
ukbm160496

# Harsap

Putra Batam
Username baru : BPR_PB
Password baru  : @08bpr08PB,

Berkat Sejati
Username : Berkat
Password : 4p4S@J@96.

# Server R&D 

- Host : 10.1.2.20
- port : 22
- username : ukabima
- password : Pa$$ip20#Ukbm6940!W0rd

Pa$$ip40#Ukbm6940!W0rd

Pa$$ip10#Ukbm6940!W0rd

username : berkat
pass : 4p4S@J@96.

username : 

# Server Baros

- host : 10.1.2.70
- anydesk : 867 135 574
- username : ukabima
- password : PSHPdBuDRC
- anydesk password : 4wdukabima

# Server Non UKBM(Teamviewer)

- host : 10.1.2.40
- password : Pa$$ip40#Ukbm6940!W0rd

# Production (UKBM)

- Host : 10.1.2.21
- password : Pa$$ip21#Ukbm6940!W0rd

# Database R&D & Baros

- username : ukabima
- password : PwdDB3bpr6940D!ukbmNew


# Gitlab Login

- Username : rizal_purnama22@yahoo.com
- Password : gitlabpassword

- Remove a local branch:
	git branch -d <branch-name>

- Remove a remote branch:
	git push origin :<branch-name>

# Setting Jam (Linux)

// disable automatic 
sudo timedatectl set-ntp 0  // 1 to Enable

sudo timedatectl set-time "2022-10-30 23:50:00"

=========

sudo date --set="2022-11-01 23:55:00.00"

# Setting jam (Windows)

date 31-08-2022
time 23:55:00

date 31-10-2022
time 23:57:00


# eBPR Login

- Username : ukabima
- Password : 4p4S@J@96.

- username : Spvr020
- password : Ukabima1.

## BS

- Username : IT2
- Password : Berkat1.

# Password Laptop

Ukabima1.

# Edit pake Vim

shift + i : untuk ngedit
setelah edit, pencet shift + : , trus ketik wq, trus enter

# Unit Testing from terminal

1. Masuk ke directory test
2. go test -v -run 'TestNamaTest' || go test -v -count=1 -timeout 999999m -run TestRunEngine

unit test & all banckmark 

	go test -v -count=1 -run=TestAsyncPointerSliceArgument -bench=.

All benchmark only

	go test -v -count=1 -run=TestGaAda -bench=.

One benchmark only

	go test -v -count=1 -run=TestGaAda -bench=BenchmarkAsyncPointerSliceArgument

# Dump a Table (schema only)

## Local
pg_dump -h localhost -p 5432 -U rizal -s -t temporary.tmp_laporan_engine_pencairan_kredit_sebulan berkat > D:\rizal\db-backups\tmp_laporan_engine_pencairan_kredit_sebulan.sql

## Production UKBM
pg_dump -h localhost -p 5432 -U ukabima -s -t temporary.tmp_laporan_engine_pencairan_kredit_sebulan bpr > /home/ukabima/rizal/dumped-schema/tmp_laporan_engine_pencairan_kredit_sebulan.sql

# Dump a Table

## Local
pg_dump -h localhost -p 5432 -U rizal --format plain -t public.m_setting_engine_golang bpr > D:\rizal\db-backups\m_setting_engine_golang.backup

## Production UKBM
pg_dump -h localhost -p 5432 -U ukabima --format plain -t public.m_setting_engine_golang bpr > /home/ukabima/rizal/dumped-schema/m_setting_engine_golang.backup

# Restore Dumped Schema

psql -d tab_bs2_go -h localhost -p 5432 -U rizal < D:\rizal\db-backups\tmp_laporan_nominatif_kredit.sql

# Reset db

SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'berkat' AND pid <> pg_backend_pid();
DROP DATABASE if EXISTS berkat;
CREATE DATABASE berkat;

# Restore DB

Baros 

pg_restore --host localhost --port 5432 --username "ukabima" -W --dbname "ukbm_2022110110" --role "ukabima" --verbose "/home/ukabima/backups/db_backups/UKBM/db_backup_ebpr__20221101100001__.backup"

Local (pass : Navicate1.)

pg_restore --host localhost --port 5432 --username "rizal" -W --dbname "go_1" --role "rizal" --verbose "D:\rizal\db-backups\db_backup_ebpr__20221101070001__.backup"


# Proses Deploy Engine Golang

- Masuk ke project directory.

- Cek db config (wajib)
	nano config/db_config.go

- Build Project
	cd cmd/engine/
	go build go-engine-backend.go

- Buat Service file (kalo belum ada) dir:/etc/systemd/system

	nama-nama service production : engine_ukbm, engine_bs, engine_pb, engine_abs
	nama-nama service testing : golangapp, golangappcahyo

- Restart Service
	service <nama-service> restart

nano Workspace/rizal/golang/config/db_config.go

cd Workspace/rizal/golang/cmd/engine | go build go-engine-backend.go

# Proses Deploy Java di Linux

1. Ubah DB Setting

	- App
	nano src/main/java/com/tdi/configuration/HibernateConfiguration.java

	- Engine
	nano src/main/webapp/WEB-INF/applicationContext.xml

2. Clean & Build
	mvn clean package

3. Pastikan tidak ada file .War yang sama pada tomcat/webapps/	
	cd /var/lib/tomcat8/webapps/

	rm -rf rizal_engine rizal_engine.war

4. Copy War file to tomcat/webapp

	- App
	cp target/tdi-bpr.war /var/lib/tomcat8/webapps/<nama-web>.war

	- Engine
	cp target/engine.war /var/lib/tomcat8/webapps/rizal_engine.war

5. Cek Proses Deploy
	tail -f /var/lib/tomcat8/logs/catalina.out

6. Akses Aplikasi Web di browser

# Rounding Mode in Java

https://docs.oracle.com/javase/7/docs/api/java/math/RoundingMode.html


# Terminate Port on Windows

netstat -ano | findstr :8080
taskkill /PID  /F

# Setting golang env linux
https://www.cs.purdue.edu/homes/bb/cs348/www-S08/unix_path.html                                                                                                                 

# Manage Multiple Session linux
https://ostechnix.com/screen-command-examples-to-manage-multiple-terminal-sessions/


# Error handling
- Perhatiakn perulangan di dalamn perulangan di dalam proses engine service, apakah jika terjadi error akan continue atau break
	> coba yang proses AutodebetSetoranRutinBerjanga gw bikin break

# Auntenticate Github
https://stackoverflow.com/questions/19660744/git-push-permission-denied-public-key


# Go Unit Testing
- melakukan testing dengan menghapus cache terlebih dahulu.
	go clean -testcache && go test -v -run "TestGoroutineLagi"

# Money & Accounting Data Types

Please do not use float64 to count money. Floats can have errors when you perform operations on them. Using big.Rat (< Go 1.5) or big.Float (>= Go 1.5) is highly recommended.

# Create a Sysmtemd Service 

1. Create New file.service in /etc/systemd/system

2. Write code below :

	[Unit]
	Description=MyApp Go Service
	After=network.target

	[Service]
	User=root
	Group=root
	ExecStart=/home/ukabima/workspace/rizal/golang/cmd/engine/go-engine-backend
	Restart=on-failure
	RestartSec=10
	StandardOutput=syslog
	StandardError=syslog
	SyslogIdentifier=appgoservice

	[Install]
	WantedBy=multi-user.target

3. Restart deamon
	
	sudo systemctl daemon-reload

4. Start your service
	
	sudo systemctl start your-service.service

5. To check the status of your service

	sudo systemctl status example.service

6. To enable your service on every reboot
	
	sudo systemctl enable example.service

7. To disable your service on every reboot

	sudo systemctl disable example.service

# Makefile Notes

	Error :

		makefile:2: *** missing separator.  Stop.
		makefile:2: *** multiple target patterns. Stop.

	Solution :

		On VS Code, just click the "Space: 4" on the downright corner and change it to "Convert Indentation to Tabs" when editing your Makefile.

# Golang Migrate

	Instalation : https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

	3 Steps to Migrate :

		1. Create migration

			migrate create -ext sql -dir <directory> -seq <migration-name>

		2. Define sql syntax on migrate up and migrate down (migrate down is redo all of migrate up)

		3. Run Migration

			migrate -path <path-directory> -database "postgresql://<username>:<password>@localhost:5432/<db_name>" -verbose up

			migrate -path <path-directory> -database "postgresql://<username>:<password>@localhost:5432/<db_name>" -verbose down

# SQLC

Tahapan dalam implementasi SQLC

1. Initial sqlc

	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc init

2. Define configurasi di sqlc.yml

3. Buat minimal satu sql query di path yang sudah ditentukan dalam sqlc.yml

3. Generate & Build sqlc

	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

## Error & Solution :

	error parsing queries: no queries contained in paths /src/db/query

	Tidak ada sql query di path yang sudah ditentukan

# Free Bootstrap template

	- https://demo.getstisla.com/forms-advanced-form.html


# Belajar Docker

- Run docker image as container with published port, and defined name

	docker run -d -p <host-port>:<container-port> --name <container-name> <image-name>

- Show list of running container 

	 docker ps

- Show list of all avaiable container

	docker ps -a

- Remove containers

	docker rm <container-name> <container-name>

- Create Volume
	
	docker volume create <volume-name>

- Create Network

	docker network create -d <network-driver> <network-name>

- Set Environment Varriable on Container

	docker container create --name <container-name> --publish <host-port>:<container-port> --env <env-name>=<value> <image-name>:<tag>

- Set Resource Limit

	docker container create --memory 500m --cpus 0.5 --name <container-name> --publish <host-port>:<container-port> --env <env-name>=<value> <image-name>:<tag>

- Show container statistic

	docker container stats

- Bind Mounts

	--mount "type=bind,source=<local-directory>,destination=<container-data-dir>"

- Bind Volume

	--mount "type=volume,source=<volume-name>,destination=<container-data-dir>"

Note : Nilai destination diisi dengan lokasi dimana data disimpan di dalam containernya (lihat di dokumentasi hub.docker masing" image)

## Docker Volume

- Show Volume List

	docker volume ls

- Create volume

	docker volume create <volume-name>

- Remove volume

	docker voulme rm <volume-name>

- Backup Volume

	Tahapan Backup 

	(step by step)

	1. Matikan container yang menggunakan volume yang akan dibackup

	2. Buat container baru dengan dua mount, satu mount untuk volume yang ingin kita backup, satu lagi untuk bind mount ke host untuk menyimpan data backupan volume.

	docker container create --name <container-name> --mount "type=bind,source=<folder-backup-host>,destination=<folder-backup-container>" --mount "type=volume,source=<volume-name>,destination=<container-dir>" <docker-image>

	3. lakukan backup dengan melakukan archive isi volume kemudian disimpan ke host directory

	tar cvf /backup/backup-mongo_1.tar.gz /data

	4. delete container yang digunakan untuk melakukan backup

	5. jalankan kembali container yang sudah dibackup


	(simple)

	1. Matikan container yang menggunakan volume yang akan dibackup

	2. Jalankan container bakcup

	docker container create --name <container-name> --mount "type=bind,source=<folder-backup-host>,destination=<folder-backup-container>" --mount "type=volume,source=<volume-name>,destination=<container-dir>" <docker-image> tar cvf /backup/backup-mongo_2.tar.gz /data

	3. jalankan kembali container yang sudah dibackup

- Restore Volume

	1. buat volume baru

	2. docker container run --rm --name ubuntubackup --mount "type=bind,source=D:\rizal\Go-Projects\belajar-docker\backups,destination=/backup" --mount "type=volume,source=mongorestore,destination=/data" ubuntu:latest bash -c "cd /data && tar xvf /backup/backup-mongo_1.tar.gz --strip 1"

	3. buat container yang menggunakan volume baru yang sudah direstore

## Docker Network

- Create network

	docker network create -d <bridge/none/host> <network-name>
	ex: docker network create -d bridge contohnetwork

- Remove network

	docker network rm <network-name>
	ex: docker network rm contohnetwork

Contoh case buat 2 container yang saling terhubung (mongodb & mongo-express)

	1. buat network
		docker network create -d bridge mongodbnetwork

	2. buat container mongodb (tanpa publish port)
		docker container create --memory 500m --cpus 0.5 --name mongodb --network mongodbnetwork --env MONGO_INITDB_ROOT_USERNAME=rizal --env MONGO_INITDB_ROOT_PASSWORD=rizal mongo:latest

	3. pull mongo-express image
		docker pull mongo-express:latest

	4. buat container mongo-express (publish port)
		docker container create --memory 500m --cpus 0.5 --name mongo-express --network mongodbnetwork --publish 8081:8081 --env ME_CONFIG_MONGODB_ADMINUSERNAME=rizal --env ME_CONFIG_MONGODB_ADMINPASSWORD=rizal --env ME_CONFIG_MONGODB_URL=mongodb://rizal:rizal@mongodb:27017/ mongo-express:latest

	5. akses mongo-ekspres dari web

- Menghapus container dari network

	docker network disconnect <nama-network> <nama-container>
	ex: docker network disconnect mongodbnetwork mongodb

- Menambahkan container kedalam network

	docker network connect <network-name> <container->
	ex: docker network connect mongodbnetwork mongodb

## Docker inspect

	Digunakan untuk menampilkan informasi detail image, volume, network, atau container.

	docker container inspect <container-name>
	docker network inspect <network-name>
	docker volume inspect <volume-name>

	ex: docker container inspect mongodb

## Docker prune

Digunakan untuk membersihkan container, image, dan volume yang sudah tidak terpakai

- menghapus semua container yang sudah stop

	docker container prune

- menghapus semua image yang tidak digunakan

	docker image prune

- menghapus semua network yang tidak digunakan

	docker network prune

- menghapus semua volume yang tidak digunakan

	docker volume prune

- menghapus semua komponen yang tidak digunakan (kecuali volume)

	docker system prune

## Dockerhub

- login (biar bisa melakukan push image ke docker hub repository)

	docker login -u <username> -p <password>
	ex: docker login -u rizalpurnama -p a7xfoldizm

- push 

	docker push <username>/<image>:<tag>
	ex: docker push rizalpurnama/from:1.0.0

- logout

	docker logout

## Docker notes

- Di image alpine untuk install aplikasi bisa pakai apk

	ex (install curl): RUN apk --no-cache add curl

# Dockerfile

File yang berisi instruksi" untuk menjalankan suatu program yang kemudian menjadi sebuah docker image

- Format penamaan file : Dockerfile
- Format Instruksi : UPPPER CASE
- Nama image saat melakukan build bisa disesuaikan dengan nama docker hub
	-t <dockerhub-account>/<image-name>:<tag>
	ex: -t khannedy/sample:1.0

## From

Digunakan saat build stage (proses pembuatan) sebagai layer dasar dari docker image yang akan dibuat

	FROM <image>:<tag>
	ex: FROM alpine:3

## RUN

Digunakan untuk menjalankan operasi saat proses pembuatan image (build stage)

	RUN command
	RUN [command, parameter]

	ex: RUN mkdir hello

## Command

Instruksi yang digunkanan untuk menjalankan suatu operasi saat sebuah image sudah dijadikan container dan container tersebut dijalankan.

Hanya bisa satu, jika ada lebih dari satu maka hanya menjalankan yang terakhir

	CMD command param param
	CMD ["executable", "param", "param"]
	CMD ["param", "param"] (menggunakan executable entry point)
	ex: CMD cat "hello/world.txt"

## Label

Digunakan untuk memberikan informasi (metadata) pada docker image seperti informasi lisensi, author, company, dll.

	LABEL key value
	LABEL key="value"
	LABEL key1="value1" key2="value2"
	ex: LABEL company="rizalpurnama" website="https://www.rizalpurnama.com"

## Add

Digunakan untuk menambahkan data dari direktori atau url kedalam image.

Jika file yang akan ditambahkan berupa archive, maka perintah add juga akan secara otomatis mengekstrak.

Perintah add dapat menggunakan Match pattern pada golang https://pkg.go.dev/path/filepath

	ADD <source> <destination-path>
	ex: ADD text/world.txt hello

## Copy

Berfungsi seperti ADD akan tetapi tidak dapat menerima source berupa url serta tidak bisa secara otomatis melakukan ekstrak archive file. Namun perintah COPY lebih direkomendasikan jika hanya ingin melakukan copy file kedalam docker image.

	COPY <source> <destination-path>
	ex: COPY text/world.txt hello

## Ignore

Berfungsi mirip seperti gitignore dalam proses ADD atau COPY
	
	#ignore a file
	text/world.txt

	#ignore a folder
	text/temp

	#ignore file with pattern
	text/*.log

## Expose

Berguna untuk memberikan informasi bahwa image ini mengekspose port dan protocol tertentu sehingga dapat diketahui saat akan membuat container dari image tersebut.

	EXPOSE <port>/<protocol>
	ex: EXPOSE 8089

## Volume

Berguna untuk secara otomatis membuat sebuah volume dan menggunakan volume tersebut ketika membuat container membuat image
	
	VOLUME /dir-name

	ENV APP_DATA=/data
	VOLUME ${APP_DATA}

## Workdir

Berfungsi untuk menentukan $HOME atau working directory untuk menjalankan perintah" RUN, CMD, ENTRYPOINT, COPY dan ADD.

Jika working directory belum ada, makan aka secara otomatis dibuatkan.

Jika working directory berupa relative path maka akan dimasukkan ke directory sebelumnya

	WORKDIR /app (absolute path)
	WORKDIR bolang (relative path, akan menjadi /app/bolang)
	WORKDIR /home/work (absolute path)

## User

Digunakan untuk mendefinisikan user yang menjalankan perintah" di golang (jika tidak mau pakai root).

User yang akan dipakai harus sudah ada, jika tidak maka harus baut user terlebih dahulu. Proses pembuatan user dan group bisa dilakukan di dockerfile sebelum menggunakan perintah USER.

	USER user-name
	ex: USER myname

## Argument

Mirip seperti ENV tapi hanya dapat digunakan pada saat proses build, dan value.nya dapat secara dinamis ditentukan saat melakukan build image dengan perintah --build-arg <arg-name>=<value>

Jika ingin mengakses nilai dari argument pada saat runtime, maka nilai tersebut harus disimpan dalam env terlebih dahulu kemudian program dapat mengakses nilai env tersebut.

	ARG name-argument
	ARG name-argument=default-value

## Healthchek

Mekanisme yang melakukan pengecekan pada suatu container yang sedang berjalan secara berkala.

	HEALTHCHECK NON #healthchek disabled (DEFAULT)
	HEALTHCHECK [OPTION] CMD command

	OPTION:
		> --interval=duration (default: 30s) #interval pengecekan
		> --timeout=duration (default: 30s) #timeout yang diberikan ketika saat menjalankan pengecekan tak kunjung ada respon
		> --start-period=duration (default: 0s) #interval pengecekan pertama dari awal dijalankannya container
		> --retries=N (default: 3) #percobaan yang diberikan saat pengecekan tidak berhasil sebelum dinyatakan sebagai unhealth

	ex: HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=3 CMD curl -f http://localhost:8089/health

## Entrypoint

Digunakan sebagai entrypoint atau executable yang akan dieksekusi saat container dijalankan. jika menggunakan entrypoint maka instruksi berisi parameter dibutuhkan oleh entrypoint

	ENTRYPOINT ["executable", "param1", "param2"]
	ENTRYPOINT executable param1 param2
	ex: 
		ENTRYPOINT [ "go", "run" ]
		CMD ["app/main.go"]

## Multi Stage Build

Mekanisme melakukan lebih dari satu build stage yang berguna untuk mengoptimalisasi penggunakan base image. Dilakukan dengan menggunakan beberapa base image sesuai dengan kebutuhannya (compile, dll) kemudian menggunakan base image dengan ukuran kecil untuk base image akhir, dengan begitu ketika melakukan build image ukurannya akan kecil sesuai dengan base image yang terakhir.

Base image selain yang terakhir harus diberi nama (alias)

	FROM golang:1.19.3-alpine as compiler #base image untuk compile project golang

	WORKDIR /app

	COPY main.go .

	RUN go build -o ./main ./main.go


	FROM alpine:3	#base image yang digunakan untuk menjalankan file hasil compile dari builder diatas. juga sebagai base image akhir 

	WORKDIR /app

	COPY --from=compiler /app/main .

	ENTRYPOINT ./main

# Docker Compose

- Merupakan tools untuk otomatisasi pembuatan dan konfigurasi satu atau beberapa kontainer, volume, network secara sekaligus

- konfigurasi disimpan didalam yaml file, defaultnya : docker-compose.yml

- Versi docker compose file dapat dilihat di : https://docs.docker.com/compose/compose-file/compose-versioning/

- Docker compose file v3 : https://docs.docker.com/compose/compose-file/compose-file-v3/

## Create

- Setelah mendefinisikan configurasi container pada docker-compose.yaml file, ketikkan perintah untuk membuat container. Lokasi eksekusi perintah harus sama dengan letak docker-compose.yaml berada.

	docker compose create

- Jika menjalankan create untuk yang kedua kali, maka docker tidak akan membuat container yang sudah terbuat sebelumnya.

## Start

- Digunakan untuk menjalankan semua container yang ada di docker-compose.yaml (containernya harus sudah dibuat terlebih dahulu)

	docker compose start

## Container List

- Digunakan untuk menampilkan semua container yang terdapat pada spesific docker-compose.yaml (bukan semua container yang ada di docker)

	docker compose ps

## Stop

- Digunakan untuk mematikan semua container yang terdapat pada spesific docker-compose.yaml
	
	docker compose stop

## Down

- Digunakan untuk menghapus semua container, network, dan volume yang ada pada spesific docker-compose.yaml termasuk juga container yang sedang running (secara otomatis akan distop terlebih dahulu sebelum dihapus).

	docker compose down

## Up

- Digunakan untuk membuat semua container, volume, dan network yang ada di specific docker-compose.yaml sekaligus menjalankannya.

- Tambahkan -d supaya semua proses dilakukan di background, bukan di terminal yang sedang dipakai.

	docker compose up -d

## Project Name

- Project name pada docker compose akan secara otomatis disamakan dengan nama folder dimana file docker-compose.yaml disimpan.

- Untuk melihat semua docker compose project yang sedang berjalan beserta lokasi config file.nya:

	docker compose ls

## Service

- Configurasi container dalam docker compose didefiniskan di bagian services.
- Dapat mendefinisikan satu atau lebih container.
- Docs: https://docs.docker.com/compose/compose-file/compose-file-v3/#service-configuration-reference

## Ports

Melakukan expose port yang ada di container ke host port.

Short Syntax :

	ports:
		- "8081":"80"		#<host-port>:<container-port>

Long Syntax :

	ports:
		- protocol: tcp 	# protocol yang digunakan (default: tcp)
		- published: 8080	# host-port
		- target: 80		# container-port

## Environment Variable

Memberikan environment variable pada configurasi container.

	environment:
		key:value

Untuk environment apa saja yang ada di image bisa di lihat di docker hub image terkait.

## Bind Mount

Mengintegrasikan local directory dengan container

Short Syntax:

	<container-config>
	volumes:
		- "./data-mongo1:/data/db"		# <local-path>:<container-path> (path data yang di container bisa dilihat di docker hub docs)

Long Syntax:

	<container-config>
	volumes:
		- type: bind
		  source: "./data-mongo2"
		  target: "/data/db"
		  read_only: false

## Volumes

Mengintegrasikan container dengan volume

Short Syntax:
	
		<container-config>
		volumes:
			- "data-mongo1:/data/db"		# <volume-name>:<container-path> (path data yang di container bisa dilihat di docker hub docs)

Long Syntax:

		<container-config>
		volumes:
			- type: volume
			  source: data-mongo2			# volume-name
			  target: "/data/db"
			  read_only: false

	volumes:
	  mongo-data1:
	    name: mongo-data1
	  mongo-data2:
	    name: mongo-data2

Tidak seperti container dan network, Volume tidak ikut dihapus ketika dilakukan docker compose down.

## Network

Saat menggunakan docker compose, sebenarnya network secara otomatis akan dibuatkan. tapi bisa juga membuat network dengan nama sesuai yang dikehendaki.

	networks:
	  network_example:				#key
	    name: network_example
	    driver: bridge

## Depends On

Menaknisme yang digunakan untuk menbuat container secara berurutan dalam docker compose. (secara default docker compose akan membuat semua container secara bersamaan sekaligus)

	services:
		<container-config>
		depends_on:
			- mongo

## Resource Limit

Mekanisme memberikan batas minimum dan maksimum resource (CPU atau Memory) pada tiap container


	services:
		<container-config>
		deploy:
			resources:
				reservations:
					cpus: "0.25"
					memory: 50M
				limits:
					cpus: "0.5"
					memory: 100M

## Build

Dalam docker compose, selain membuat container dari docker image yang sudah ada, kita juga bisa membuat sebuah container dari image yang dibuat dari Dockerfile kita sendiri. yaitu dengan menggunakan atribut build.

	services:
	  	<container-configs>
	    build:
	      context: <dockerfile-path>
	      dockerfile: <dockerfile-name> (default is Dockerfile)
	    image: <image:tag>  (jika tidak ingin dibautkan dengan nama random)

Jika ada perubahan dari image yang dipakai, maka container yang sedang berjalan tidak akan secara langsung terupdate. harus dibuild & run ulang.(tidak hot reload)

## Health Check

Selain di dockerfile, health check juga dapat di docker compose dengan mendefinisikan konfigurasi healthcecknya

	services:
	  	<container-configs>
	  	healthcheck:
	  		test: ["command", "to", "do"]
	  		interval: 5s
	  		timeout: 5s
	  		retries: 3
	  		start_period: 5s

## Extends Service

Digunakan untuk menangani perbedaan konfigurasi atau environement variable.
Membuat satu docker-compose.yaml file untuk konfigurasi yang umum, kemudian membuat beberapa yaml file yang lain untuk masing" konfigurasi.
kemudian ketika akan menjalankan docker compose, panggil docker compose yang umum diikuti dengan yaml file kedua sesuai dengan yang diinginkan, masing" menggunakan -f .

contoh : docker compose -f docker-compose.yaml -f dev.yml up -d


# MongoDB

## Database

Tidak seperti mysql/postgresql, untuk menggunakan database mongodb dapat langsung menggunakan keyword use <nama-db> tanpa perlu membuat databasenya terlebih dahulu.

	use <db-name>

Show list databases

	show databases

https://www.mongodb.com/docs/mongodb-shell/reference/methods/#std-label-mdb-shell-methods

## Data Model

Embedded:
	Digunakan jika antar document saling tergantung, embedded document selalu dibutuhkan ketika mengambil data document, tidak bisa melakukan perubahan secara langsung pada embedded document

	contoh :

	{
		_id: <ObjectId1>
		username: "123Abc"
		contact: {
			phone: "0829191992",
			email: "sha@mail.com"
		 },
		access: {
			level: "5",
			group: "dev"
		}
	}

Refference:
	Digunakan jika antar document tidak saling bergantung (kebalikan dari embedded). mirip seperti relasional database, tapi untuk mendapatkan datanya tidak bisa pakai join, tapi harus melakukan query ulang ke document yang diinginkan.

	user document:

	{
		_id: <ObjectId1>
		username: "123Abc"
	}

	contact document:

	{
		_id: <ObjectId2>
		user_id: <ObjectId1>
		phone: "0829191992",
		email: "sha@mail.com"
	}

	access document:

	{
		_id: <ObjectId3>
		user_id: <ObjectId1>
		level: "5",
		group: "dev"
	}

## Insert

Field _id secara default akan dibuatkan ObjectId yang berupa random string, akan tetapi bisa juga diisi oleh data apasaja yang berdungsi sebagai primary key. atau bisa dengan secara explisit menyebutkan new ObjectId()

Secara default, nilai angka akan dibaca sebagai double. jika ingin spesifik membuatnya menjadi long (64 integer) yaitu :

	new NumberLong(1000)

int : new NumberInt(100)

https://www.mongodb.com/docs/manual/reference/operator/query/
https://www.mongodb.com/docs/manual/reference/method/
https://www.mongodb.com/docs/manual/reference/

## Mongo <> Golang
https://www.mongodb.com/docs/drivers/go/current/
https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.11.0

# Golang WEB

## Server

Komponen yang berfungsi sebagai web server dan direpresentasikan dalam struct bernama Server

	func TestServer(t *testing.T) {
		server := http.Server{
			Addr: "localhost:8181",
		}

		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}

## Handler

Komponen di dalam server yang menangani request yang masuk pada sebuah web server, tanpa handler web server tidak bisa menangani request dari web browser/client.

	func TestHandler(t *testing.T) {
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello World!")
		}

		server := http.Server{
			Addr:    "localhost:8081",
			Handler: handler,
		}

		server.ListenAndServe()
	}


### HTTP response status codes

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status


# Docker Events

Memonitoring semua event yang terjadi pada docker

	docker events -f 'container=<nama-container>'
	ex: docker events -f 'container=mongo-express-example'

# Docker commands

- Build Image : docker build --tag ebpr-engine:1.0 .
- Remove Image : docker image rm -f image-id
- Create container :	docker container create --name my-container my-image:1.0
						docker run -d --network=host --name=ukabima ebpr-engine:1.0
- Run container : docker container start my-container
- Remove container : docker container rm my-container
-  : docker container create --network=host --name my-container my-image

# Build Golang Back End project from scracth

- Create database schema using DBML on dbdiagram.io and than export it to postgreSQL.
- Generate db schema from exported dbdiagram file.
- Add uuid generator extension (if not exist) Error : function uuid_generate_v4() does not exist

	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

- Create new project directory & initialize go module
- Create makefile and define syntax to run postgres as a docker container and than create db and drop db.
- Run db migration with golang migrate
- 

	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5434/technical-test?sslmode=disable" -verbose up

Articles :
	- Unix file system permission 
		https://www.guru99.com/file-permissions.html
		https://en.wikipedia.org/wiki/File-system_permissions#:~:text=0666,read%3B%20others%20have%20no%20permissions

### menjalankan engine as a service

	- Jika merubah tanggal, service harus direstart
	- merubah isi service harus : daemon-reload, kemudian restart servicenya.

### Yang perlu kasih default value di table
	- acc_journal
	- acc_journal_detail
	- m_nomor_kode_jurnal
	- m_nomor_vourcher


### Saat melakukan test, pastikan dulu settingan sudah benar
### Saat membuat model, dahulukan foreign key field.
### Saat akan membuat function, find usage terlebih dahulu. untuk menyesuaikan parameter.
### Saat Membuat Function Harus buat log untuk tracking start & stop
### Setting golang env linux
https://www.cs.purdue.edu/homes/bb/cs348/www-S08/unix_path.html                                                                                                                 

### Manage Multiple Session linux
https://ostechnix.com/screen-command-examples-to-manage-multiple-terminal-sessions/


### Error handling
- Perhatiakn perulangan di dalamn perulangan di dalam proses engine service, apakah jika terjadi error akan continue atau break
	> coba yang proses AutodebetSetoranRutinBerjanga gw bikin break

### Auntenticate Github
https://stackoverflow.com/questions/19660744/git-push-permission-denied-public-key


### Go Unit Testing

- melakukan testing dengan menghapus cache terlebih dahulu.
	go clean -testcache && go test -v -run "TestGoroutineLagi"

### Convert JSON to YAML

https://www.json2yaml.com/

### Create Website
squarespace

### Regex
https://regexr.com