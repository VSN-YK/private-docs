package draw.com;

import java.awt.Container;

import javax.swing.JButton;
import javax.swing.JFrame;

public class BaseFrame extends JFrame {
	protected Container c;

	public BaseFrame(String title) {
		setTitle(title);
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		c = getContentPane();
		setPane(c);
	}
	private void setPane(Container c) {
		setSize(400,300);

		DrawAppCore appCore = new DrawAppCore();
		DrawPanel drawPanel = new DrawPanel(appCore);
		appCore.setDrawPanel(drawPanel);

		DrawMouseListener ml = new DrawMouseListener(appCore);
		drawPanel.addMouseListener(ml);
		drawPanel.addMouseMotionListener(ml);

		SelectListener sl = new SelectListener(appCore);
		JButton rectBtn = new JButton("line");
		JButton custom = new JButton("custom");
		rectBtn.addActionListener(sl);
		custom.addActionListener(sl);
		rectBtn.setActionCommand("line");
		rectBtn.setActionCommand("custom");
		drawPanel.add(rectBtn);
		drawPanel.add(custom);

		c.add(drawPanel);
		setVisible(true);
	}
}
