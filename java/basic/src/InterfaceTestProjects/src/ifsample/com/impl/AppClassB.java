package ifsample.com.impl;

import ifsample.com.base.ImplBase;

public class AppClassB implements ImplBase {

	private String appName = "B";

	@Override
	public String getAppVersion() {
		// TODO 自動生成されたメソッド・スタブ
		return this.appName;
	}

}
