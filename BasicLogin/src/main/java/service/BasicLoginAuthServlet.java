package service;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;

import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import model.dao.Dao;
import model.dao.DaoFactory;
import model.dao.LoginDao;
import model.dto.LoginDto;

@WebServlet("/login")
public class BasicLoginAuthServlet extends HttpServlet {
	// スーパクラスであるHttpServletのコンストラクタを呼び出す。
	LoginDao loginDao;

	public BasicLoginAuthServlet() {
		super();
	}

	// TODO 検証ようのため、GETでリクエストを送る。のちにPOSTでreceiveするようにする。
	@Override
	protected void doGet(HttpServletRequest req, HttpServletResponse res) {
		try {

			loginDao = DaoFactory.createLoginDao();

			if(getBasicAuth(loginDao.selectAll(),"L001","Node.js")) {
				// ログイン成功ページへフォワード
				System.out.println("認証に成功しました。");
			}else {
				// ログイン失敗ページへリダイレクト
				System.out.println("認証失敗しました。");
			}
		} catch (SQLException e) {
			e.printStackTrace();
		}finally {
				Dao dao = loginDao;
				try {
					dao.getConnection().close();
				} catch (SQLException e) {
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
	private boolean getBasicAuth(ResultSet rs, String id, String name) throws SQLException {
		while (rs.next()) {
			if ((id.equals(rs.getString("id"))) && (name.equals(rs.getString("name")))) {
				return true;
			}
		}
		return false;
	}

	/**
	 * BASIC認証を行うビジネスメソッド(オーバロード)
	 * @param loginList
	 * @param id
	 * @param name
	 * @return boolean
	 * @throws SQLException
	 */

	private boolean getBasicAuth(List<LoginDto> loginList, String id, String name) throws SQLException {
		for (LoginDto l : loginList) {
			if ((l.getId().equals(id)) && (l.getName().equals(name))) {
				return true;
			}
		}
		return false;
	}
}
