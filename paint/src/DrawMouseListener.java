package draw.com;

import java.awt.event.MouseEvent;

import javax.swing.SwingUtilities;
import javax.swing.event.MouseInputAdapter;

public class DrawMouseListener extends MouseInputAdapter {
	protected DrawAppCore appCore;
	protected int dragStartPointX, dragStartPointY;

	public DrawMouseListener(DrawAppCore app) {
		appCore = app;
	}

	@Override
	public void mousePressed(MouseEvent e) {
		dragStartPointX = e.getX();
		dragStartPointY = e.getY();

		if(SwingUtilities.isRightMouseButton(e) == true) {
			appCore.eachUndo();
		}else if(SwingUtilities.isLeftMouseButton(e) == true) {
			appCore.generateFigure(dragStartPointX, dragStartPointY);
		}
	}

	@Override
	public void mouseReleased(MouseEvent e) {
		appCore.reDrawShape(dragStartPointX, dragStartPointY, e.getX(), e.getY());
		
	}

	@Override
	public void mouseDragged(MouseEvent e) {
		appCore.reDrawShape(dragStartPointX, dragStartPointY, e.getX(), e.getY());
	}
}

