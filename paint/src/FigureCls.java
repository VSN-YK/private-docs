package draw.com;

import java.awt.Color;
import java.awt.Graphics;
import java.awt.Rectangle;

public abstract class FigureCls  {

	protected Rectangle legendRect = new Rectangle();
	protected Color color;

	/**
	 * @param x
	 * @param y
	 * @param w
	 * @param h
	 * @param c
	 */
	public FigureCls(int x , int y , int w, int h, Color c) {
		setLocation(x, y);
		setSize(w, h);
		color = c;
	}

	public void setLocation(int x , int y) {
		legendRect.setLocation(x,y);
	}

	public void setSize(int w, int h) {
		legendRect.setSize(w,h);
	}

	/**
	 *
	 * @param g
	 */
	protected abstract void paint(Graphics g) ;

	/**
	 *
	 * @param x0
	 * @param y0
	 * @param x1
	 * @param y1
	 */
	protected abstract void reDraw(int x0 , int y0, int x1, int y1);

}
