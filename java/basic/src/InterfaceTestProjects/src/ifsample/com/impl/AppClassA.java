package ifsample.com.impl;

import ifsample.com.base.ImplBase;

public class AppClassA implements ImplBase {

	private String appName = "A";

	@Override
	public String getAppVersion() {
		// TODO 自動生成されたメソッド・スタブ
		return this.appName;
	}
}
