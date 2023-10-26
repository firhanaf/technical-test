-- Soal 1
create database db_soal1;
use db_soal1;

create table mahasiswa(
nim bigint unsigned primary key auto_increment not null unique,
nama varchar(100),
kota varchar(100),
tahun_masuk year
);

create table dosen(
nip bigint unsigned primary key not null unique,
nama varchar(100),
kota varchar(100)
);

create table mata_kuliah(
kode varchar(25) primary key not null unique,
nama varchar(100),
nip bigint unsigned,
sks tinyint,
constraint dosen_matakuliah foreign key(nip) references dosen(nip)
);

show tables;
select * from mahasiswa;
select * from dosen;
select * from mata_kuliah;