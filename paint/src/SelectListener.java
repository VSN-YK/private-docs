package draw.com;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class SelectListener implements ActionListener {
	DrawAppCore appCore;

	public SelectListener(DrawAppCore app) {
		appCore = app;
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		String cmd = e.getActionCommand();
		switch (cmd) {
		case ConstantFigureUtil.LINE_FIGURE:
			appCore.updateActvKey(ConstantFigureUtil.LINE_FIGURE);
			break;
		case ConstantFigureUtil.CUSTOM_COLOR:
			appCore.updateRGBColor(255, 0, 255);
			break;
		case ConstantFigureUtil.EXIT_APP:

			System.exit(0);
		default:
			break;
		}
	}

}
