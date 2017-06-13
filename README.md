golang 语言实现mysql 对应mybatis mapper ,java 类，manager , service 生成。
示例工程：请看demo

示例：
tb_admin 表：
CREATE TABLE `tb_admin` (

`ID` BIGINT(20) NOT NULL AUTO_INCREMENT,

`NAME` VARCHAR(64) DEFAULT NULL,

`pass` VARCHAR(128) DEFAULT NULL,

 PRIMARY KEY (`ID`)

) ENGINE=INNODB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8

生成mapper 文件：
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >
<mapper namespace="com.masz.demo.model.Admin">

    <resultMap type="com.masz.demo.model.Admin" id="AdminMap">
        <id column="ID" jdbcType="BIGINT" property="id"/>
        <result column="NAME" jdbcType="VARCHAR" property="name"/>
        <result column="PASS" jdbcType="VARCHAR" property="pass"/>
    </resultMap>

    <sql id="Base_Column_List">
        ID,
        NAME,
        PASS
    </sql>

    <insert id="insert" parameterType="com.masz.demo.model.Admin">
        INSERT INTO TB_ADMIN(
        ID,
        NAME,
        PASS
        )
        values(
        #{id,jdbcType=BIGINT},
        #{name,jdbcType=VARCHAR},
        #{pass,jdbcType=VARCHAR}
        )
    </insert>

    <delete id="delete" parameterType="long" >
        delete from tb_admin
        where id = #{id}
    </delete>

    <update id="update" parameterType="com.masz.demo.model.Admin">
        UPDATE TB_ADMIN
        <set>
            <if test="name!=null">
            NAME=#{name,jdbcType=VARCHAR},
            </if>
            <if test="pass!=null">
            PASS=#{pass,jdbcType=VARCHAR},
            </if>
        </set> 
        where id=#{id}
    </update>

    <select id="count" parameterType="map" resultType="long">
        select count(1) from tb_admin
        <where>
            <if test="id!=null">
            and ID=#{id,jdbcType=BIGINT}
            </if>
            <if test="name!=null">
            and NAME=#{name,jdbcType=VARCHAR}
            </if>
            <if test="pass!=null">
            and PASS=#{pass,jdbcType=VARCHAR}
            </if>
        </where>
    </select>

    <select id="get" parameterType="long" resultMap="AdminMap">
        select <include refid="Base_Column_List" /> from tb_admin
        where id = #{id}
    </select>

    <select id="findList" parameterType="map" resultMap="AdminMap">
        select <include refid="Base_Column_List" /> from tb_admin
        <where>
            <if test="id!=null">
            and ID=#{id,jdbcType=BIGINT}
            </if>
            <if test="name!=null">
            and NAME=#{name,jdbcType=VARCHAR}
            </if>
            <if test="pass!=null">
            and PASS=#{pass,jdbcType=VARCHAR}
            </if>
        </where>
    </select>

    <select id="findPage" parameterType="map" resultMap="AdminMap">
        select <include refid="Base_Column_List" /> from tb_admin
        <where>
            <if test="id!=null">
            and ID=#{id,jdbcType=BIGINT}
            </if>
            <if test="name!=null">
            and NAME=#{name,jdbcType=VARCHAR}
            </if>
            <if test="pass!=null">
            and PASS=#{pass,jdbcType=VARCHAR}
            </if>
        </where>
        LIMIT #{offset},#{pageSize}
    </select>

</mapper>

生成com.masz.demo.model.Admin 类：

package com.masz.demo.model;
import com.masz.demo.base.BaseModel;
public class Admin extends BaseModel {
    private String name

    private String pass


    public String getName(){
        return this.name
    }

    public void setName(String name){
        this.name = name
    }

    public String getPass(){
        return this.pass
    }

    public void setPass(String pass){
        this.pass = pass
    }

}

生成dao：
package com.masz.demo.dao;

import com.masz.demo.model.Admin;
import com.masz.demo.dao.base.MyBatisDao;
import org.springframework.stereotype.Repository;

@Repository
public class AdminDao extends MyBatisDao<Admin> {

}
其中CRUD 方法在 MyBatisDao中已经实现，相映的基类方法，查看dome 工程

使用时依赖mysql 驱动：https://github.com/go-sql-driver/mysql
