package model.dao;

import java.sql.Connection;
import java.sql.SQLException;

import db.helper.connect.ConnectionHelper;

public class Dao {
	private Connection conn;
	
	public Connection getConnection() throws SQLException {
		ConnectionHelper helper = new ConnectionHelper().getConnect("");
		return helper.getConnection();
	}

	public Connection getConn() {
		return this.conn;
	}	
}
