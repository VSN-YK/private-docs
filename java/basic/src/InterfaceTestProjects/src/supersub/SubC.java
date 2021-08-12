package supersub;

public class SubC extends SubB {

	private String subCField;
	public SubC(String setData) {
		super(setData);
		this.subCField = setData;
	}

	public void setSubCField(String data) {
		this.subCField = data;
	}

	public String getSubCField() {
		return this.subCField;
	}

	@Override
	public String subBMethod () {
		return super.SuperAMethod();
	}
}
