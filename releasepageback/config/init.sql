create table if not  exists releaseinfo (
			id int auto_increment primary key ,
			name varchar(20) ,
			path varchar(20) ,
			type  varchar(20),
			version varchar(20) ,
			md5path varchar(20),
			md5name varchar(20) ,
			comment varchar(20) , );
		create table if not  exists buildimageconninfo (
			id int auto_increment primary key,
			envname varchar(20) ,
			type varchar(20) ,
			account  varchar(20),
			passwd varchar(20));
		create table if not  exists commentfiles
        (
            id
            int
            auto_increment
            primary
            key,
            name
            varchar
        (
            20
        ) ,
            path varchar
        (
            80
        )
            );

insert into buildimageconninfo values (id,'打镜像','ssh','root@192.168.22.227','JPsoft@2023'),
                                      (id,'打镜像','toDesk','188 586 722','JPsoft@2023'),
                                      (id,'OMP测试','ssh','root@192.168.22.225','JPsoft@2023'),
                                      (id,'DCP测试','ssh','root@192.168.22.226','JPsoft@2023'),
                                      (id,'Windows10远程','toDesk','428 178 232','JPsoft@2023'),
                                      (id,'腾讯云主机远程','toDesk','783 899 137','JPsoft@2023');
insert into commentfiles values (id ,'升级语句汇总','./update_sql.tar.gz'),(id ,'常用运维命令','./jpcmd.tar.gz');

