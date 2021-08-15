package service;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;

import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import model.dao.Dao;
import model.dao.DaoFactory;
@WebServlet("/login")
public class BasicLoginAuthServlet extends HttpServlet {
	protected Connection conn;
	// スーパクラスであるHttpServletのコンストラクタを呼び出す。
	public BasicLoginAuthServlet() {
		super();
	}

	// TODO 検証ようのため、GETでリクエストを送る。のちにPOSTでreceiveするようにする。
	@Override
	protected void doGet(HttpServletRequest req, HttpServletResponse res) {
		try {
			Dao dao = DaoFactory.createDao();
			this.conn = dao.getConnection();
			ResultSet rs = conn.createStatement().executeQuery("SELECT * FROM LANGUAGE_TBL");
			if(getBasicAuth(rs,"L001","Node.js")) {
				// ログイン成功ページへフォワード
			}else {
				// ログイン失敗ページへリダイレクト
			}
		} catch (SQLException e) {
			// TODO 自動生成された catch ブロック
			e.printStackTrace();
		}finally {
			try {
				this.conn.close();
			} catch (SQLException e) {
				// TODO 自動生成された catch ブロック
				e.printStackTrace();
			}
		}
	}

	/**
	 * BASIC認証を行うビジネスメソッド
	 * @param rs
	 * @param id
	 * @param name
	 * @return
	 * @throws SQLException
	 */
	private boolean getBasicAuth(ResultSet rs ,String id , String name) throws SQLException {
		while(rs.next()) {
			if( (id.equals(rs.getString("id"))) && (name.equals(rs.getString("name"))) ){
				return true;
			}
		}
		return false;
	}
}
