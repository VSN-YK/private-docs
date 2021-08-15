package db.helper.connect;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;

import javax.naming.InitialContext;
import javax.naming.NamingException;
import javax.sql.DataSource;

public class ConnectionHelper {
	private Connection conn;
	final static String DATA_SOURCE_ENTRY = "java:/comp/env/jdbc/sub";
	/**
	 * 
	 * @param connMethodTyp
	 * @return
	 * @throws SQLException 	
	 */
	public ConnectionHelper getConnect(String method) throws SQLException {
		if(method.equals("jdbc")) {
			
			// JDBCを利用しDBへの接続を行う
			String url = "jdbc:mysql://localhost/sub";
			String user = "sub_user";
			String pass = "sub";
			
			this.conn = DriverManager.getConnection(url,user,pass);
		}else {
			// JDNIを利用しDBへの接続を行う
			try {
				DataSource ds = (DataSource) new InitialContext().lookup(ConnectionHelper.DATA_SOURCE_ENTRY);
				this.conn = ds.getConnection();
			} catch (NamingException e) {
				e.printStackTrace();
			}
		}
		return this;
	}
	
	public Connection getConnection() {
		return this.conn;
	}

	/**
	 * 動作検証用のスタブ
	 * @param args
	 * @throws SQLException 
	 */
	public static void main(String[] args) throws SQLException {
		ConnectionHelper helper = new ConnectionHelper().getConnect("jdbc");
		ResultSet rs = helper.conn.createStatement().executeQuery("select * from LANGUAGE_TBL");
		while(rs.next()) {
			System.out.println(rs.getString("name"));
		}
	}

}
