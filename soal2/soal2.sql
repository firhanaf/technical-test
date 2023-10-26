-- soal 2

create table nilai(
id int primary key not null auto_increment unique,
nilai_uts_angka tinyint not null,
nilai_uts_huruf enum("A", "B", "C", "D", "E", "F") not null,
nilai_uas_angka tinyint not null,
nilai_uas_huruf enum("A", "B", "C", "D", "E", "F") not null,
nilai_akhir_angka tinyint not null,
nilai_akhir_huruf enum("A", "B", "C", "D", "E", "F") not null,
nim bigint unsigned not null,
kode varchar(25),
constraint mahasiswa_nilai foreign key (nim) references mahasiswa(nim),
constraint mahasiswa_matakuliah foreign key (kode) references mata_kuliah(kode)
);

create table ip_semester(
id int primary key not null auto_increment unique,
tahun_ajaran year,
semester enum("ganjil", "genap"),
ip decimal(4,2),
nim bigint unsigned not null,
constraint mahasiswa_ip foreign key (nim) references mahasiswa(nim)
);

show tables;
select * from nilai;
select * from ip_semester;