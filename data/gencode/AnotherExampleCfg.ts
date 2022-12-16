import Res from "../../common/util/Res"

export class AnotherInfo {
    public ID: number// ID
    public Name: string// 姓名
    public Age: number// 年龄
}

export class AnotherExampleConfig {
    private static instance: AnotherExampleConfig
    public static Get(): AnotherExampleConfig {
        if (AnotherExampleConfig.instance == null) {
            AnotherExampleConfig.instance = new AnotherExampleConfig()
            AnotherExampleConfig.instance.init()
        }
        return AnotherExampleConfig.instance
    }

    public AnotherSlc: Array<AnotherInfo>
    public AnotherMap: Map<number, AnotherInfo>

    private init(): void {
        this.AnotherSlc = new Array<AnotherInfo>()
        this.AnotherMap = new Map<number, AnotherInfo>()

        let jsonData = Res.get<cc.JsonAsset>("json/AnotherExample", cc.JsonAsset)
        this.AnotherSlc = jsonData.json['Another']
        this.AnotherSlc.forEach(another => {
            this.AnotherMap.set(another.ID, another)
        })

    }
}
