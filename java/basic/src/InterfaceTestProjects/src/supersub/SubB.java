package supersub;

public class SubB extends SuperA {
	private String subBField;
	private String superFeild;

	public SubB(String setData) {
		subBField = setData;
		superFeild = super.SuperAMethod();
	}

	public void setSubBField(String data) {
		this.subBField = data;
	}

	public String getSubBField() {
		return "[SubB]" + this.subBField + "\n" +  "[Super]" + superFeild + "\n";
	}

	@Override
	public String SuperAMethod() {
		return "This is Sub B";
	}

	public String subBMethod() {
		return "This is Called by Sub B";
	}
}
