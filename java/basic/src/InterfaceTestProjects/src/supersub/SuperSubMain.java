package supersub;

public class SuperSubMain {
	public static void main(String[] args) {

		SuperA superA = new SuperA();
		debugConsleLog(superA.SuperAMethod());

		SubB subB = new SubB("This is Sub B [Before Update]");
		debugConsleLog(subB.SuperAMethod());
		debugConsleLog(subB.getSubBField());

		subB.setSubBField("This is Sub B [After Update]");
		debugConsleLog(subB.getSubBField());

		SubC subC = new SubC("This is Sub C");
		debugConsleLog(subB.getSubBField());
		debugConsleLog(subC.getSubCField());
		debugConsleLog(subC.subBMethod());
	}

	private static void debugConsleLog(String debug) {
		System.out.println(debug);
	}

}


