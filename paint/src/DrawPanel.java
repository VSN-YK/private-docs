package draw.com;

import java.awt.Color;
import java.awt.Graphics;

import javax.swing.JPanel;

/**
 * 図形を描画するためのパネルを生成するクラスです。
 * @author hacknatural
 *
 */
public class DrawPanel extends JPanel {
	protected DrawAppCore appCore;

	public DrawPanel(DrawAppCore app) {
		setBackground(Color.white);
		appCore = app;
	}

	/**
	 * @see ペイントの描画を行うメソッドです。<br>
	 * 処理自体はペイントアプリケーションのコア基盤を提供する<br>
	 * DrawAppCoreクラスにデリゲートさせます。
	 */
	
	@Override
	public void paintComponent(Graphics g) {
		super.paintComponent(g);
		appCore.paintComponent(g);
	}
}
