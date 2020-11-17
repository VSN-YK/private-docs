package main;

import ifsample.com.base.ImplBase;
import ifsample.com.impl.AppClassA;
import ifsample.com.impl.AppClassB;
import ifsample.com.impl.AppClassC;

public class AppMain {

	public static void main(String[] args) {

		AppClassA appClassA = new AppClassA();
		AppClassB appClassB = new AppClassB();
		AppClassC appClassC = new AppClassC();

		CustomDebugMessenger(appClassA.getClass().getName(),
							 appClassB.getClass().getName(),
							 appClassC.getClass().getName()
		);

		AppVersionResolver(appClassA, appClassB, appClassC);
	}

	/**
	 *
	 * @param impl
	 */
	private static void AppVersionResolver(ImplBase... impl) {
		for (ImplBase eachImpl : impl) {
			System.out.println(eachImpl.getAppVersion());
		}
	}

	private static void CustomDebugMessenger(String... clsName) {
		for (String eachClsName : clsName) {
			System.out.println(eachClsName + "\n");
		}
	}
}
