package draw.com;

import java.awt.Color;
import java.awt.Graphics;
import java.util.ArrayDeque;
import java.util.Deque;
/**
 * @see ペイントアプリケーションのコア基盤を提供するクラスです。
 * @author hacknatural
 *
 */
public class DrawAppCore {
	protected Deque<FigureCls> figureList;
	protected FigureCls figure;
	protected Color actvColor;
	protected String actvFigureKey;
	protected DrawPanel drawPanel;

	public DrawAppCore() {
		figureList = new ArrayDeque<FigureCls>();
		figure = null;
		actvColor =  Color.BLACK;
		actvFigureKey = "line";
	}

	/**:
	 *
	 * @param dp
	 */
	public void setDrawPanel(DrawPanel dp) {
		this.drawPanel = dp;
	}

	/**
	 * JPanelを継承したサブクラスでpaintComponentをオーバライドしているメソッドが
	 * 呼び出された時にトリガーされるメソッドです。
	 * キューに保持している全ての図形オブジェクトを描画します。
	 * @param g
	 */
	public void paintComponent(Graphics g) {
		for(FigureCls f : figureList) {
			f.paint(g);
		}
	}

	/**
	 * @see クライアントからリクエストされた図形キー文字列を元に<br>
	 * 図形オブジェクトの生成を行うメソッドです。
	 * @param x
	 * @param y
	 */
	public void generateFigure(int x , int y) {
		FigureCls f = null;

		switch (actvFigureKey) {
		case ConstantFigureUtil.LINE_FIGURE:
			f = new LineFigureCls(x, y, x, y,actvColor);
			break;
		default:
			break;
		}
		figureList.add(f);
		figure = f;
		drawPanel.repaint();
	}

	/**
	 *
	 * @param x0
	 * @param y0
	 * @param x1
	 * @param y1
	 */
	public void reDrawShape(int x0, int y0, int x1, int y1) {
		if(figure != null) {
			figure.reDraw(x0, y0, x1, y1);
			drawPanel.repaint();
		}
	}

	/**
	 * @see 図形オブジェクトに基本的なフォアカラーを設定するメソッドです。
	 * @param color
	 */
	public void updateBasicColor(Color color) {
		color = actvColor;
	}

	/**
	 * @see 図形オブジェクトに任意のRGBフォアカラーを設定するメソッドです。
	 * @param color
	 * @param red
	 * @param green
	 * @param blue
	 */
	public void updateRGBColor(int red , int green , int blue) {
		 actvColor = new Color(red,green,blue);
	}

	/**
	 * クライアントが現在選択している図形のリクエスト文字列を更新する
	 * @param key
	 */
	public void updateActvKey(String key) {
		actvFigureKey = key;
	}

	/**
	 *Queの先頭要素に位置する図形オブジェクトの削除を行う
	 */
	public void eachUndo() {
		if(figureList.isEmpty()) {
			return;
		}
		figureList.remove();
	}

	/**
	 Queに格納されている図形オブジェクトを全て削除する。
	 */
	public void undoAll() {
		for(int fidx =	0; fidx < figureList.size(); fidx++) {
			this.eachUndo();
		}
	}
}
