package model.dao;

import java.sql.SQLException;

public class DaoFactory {
	public static Dao createDao() throws SQLException{
		return new Dao();
	}
	public static LoginDao createLoginDao () throws SQLException {
		return new LoginDao();
	}
}
