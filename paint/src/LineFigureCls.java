package draw.com;

import java.awt.Color;
import java.awt.Graphics;

 class LineFigureCls extends FigureCls {

	public LineFigureCls(int x, int y, int w, int h, Color c) {
		super(x, y, w, h, c);
	}

	@Override
	protected void paint(Graphics g) {
		g.setColor(color);
		g.drawLine(legendRect.x,legendRect.y,legendRect.width,legendRect.height);
	}

	@Override
	protected void reDraw(int x0, int y0, int x1, int y1) {
		setLocation(x0, y0);
		setSize(x1, y1);
	}
}
