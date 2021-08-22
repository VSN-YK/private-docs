package model.dao;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;

import model.dto.LoginDto;

public class LoginDao extends Dao {
	public LoginDao() throws SQLException {
		super();
	}
	public List<LoginDto> selectAll() throws SQLException{
		List<LoginDto> loginList =  new ArrayList<LoginDto>();
		ResultSet rs = super.conn.createStatement().executeQuery("SELECT * FROM LANGUAGE_TBL");
		while(rs.next()) {
			LoginDto loginInfo = new LoginDto();
			loginInfo.setId(rs.getString("id"));
			loginInfo.setName(rs.getString("name"));
			loginList.add(loginInfo);
		}
		return loginList;
	}
}
