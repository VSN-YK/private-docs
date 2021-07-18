package util;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class InputTestCaseHelper {

	class InnerTestCaseValue {
		private int x = 0;
		private int y = 0;

		public InnerTestCaseValue(int x, int y) {
			this.x = x;
			this.y = y;
		}

		public int getX() {
			return x;
		}

		public int getY() {
			return y;
		}

	}
	
	/**
	 * 
	 * @return ArrayList<InnerTestCaseValue>
	 */
	public ArrayList<InnerTestCaseValue> inputIntegerTestCase() {
		ArrayList<InnerTestCaseValue> array = new ArrayList<>();
		File file = new File("./src/TestCase.txt");
		Scanner scan = null;
		try {
			scan = new Scanner(file);

			scan.useDelimiter("\n");
			while (scan.hasNext()) {
				String[] inputLine = scan.next().split(",");

				int num1 = Integer.parseInt(inputLine[0]);
				int num2 = Integer.parseInt(inputLine[1].replace("\r", ""));

				array.add(new InnerTestCaseValue(num1, num2));
			}

		} catch (FileNotFoundException fnot) {
			fnot.printStackTrace();
		} catch (NumberFormatException nfe) {
			throw new IllegalArgumentException("数値の変換に失敗");
		} finally {
			scan.close();
		}

		return array;
	}

	/**
	 * 検証用のメインスタブメソッド
	 * @param args
	 */
	public static void main(String[] args) {

		InputTestCaseHelper tester = new InputTestCaseHelper();
		ArrayList<InnerTestCaseValue> vList = tester.inputIntegerTestCase();
		for (int i = 0; i < vList.size(); i++) {
			System.out.println(vList.get(i).getX() + "\t" + vList.get(i).getY());
		}

	}

}
