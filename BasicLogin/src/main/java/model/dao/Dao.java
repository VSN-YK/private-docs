package model.dao;

import java.sql.Connection;
import java.sql.SQLException;

import db.helper.connect.ConnectionHelper;

public class Dao {

	protected Connection conn;

	public Dao() throws SQLException {
		if (this.conn == null) {
			ConnectionHelper helper = new ConnectionHelper().getConnect("");
			this.conn = helper.getConnection();
		}
	}
	
	public Connection getConnection() {
		return this.conn;
	}
}
