package draw.com;

import javax.swing.JFrame;
import javax.swing.SwingUtilities;

public class DrawMain {

	public static void main(String[] args) {
		SwingUtilities.invokeLater(() -> {
		JFrame f = new BaseFrame("Draw Application");
//		JFrame f = new JFrame("Draw Application");
//		f.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
//		JPanel panel = new JPanel();
//		JButton button = new JButton("rect");
//		button.setActionCommand("rect");
//
//		DrawAppCore appCore = new DrawAppCore();
//		DrawPanel dp = new DrawPanel(appCore);
//		appCore.setDrawPanel(dp);
//		DrawMouseListener ml = new DrawMouseListener(appCore);
//		dp.addMouseListener(ml);
//		dp.addMouseMotionListener(ml);
//		panel.add(button);
//
//		SelectListener sl = new SelectListener(appCore);
//		button.addActionListener(sl);
//
//		f.getContentPane().add(dp, BorderLayout.CENTER);
//		f.getContentPane().add(panel, BorderLayout.NORTH);
//
//		f.setSize(400,300);
//		f.setVisible(true);
		});
	}
}
