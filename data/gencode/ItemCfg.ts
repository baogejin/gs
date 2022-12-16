import Res from "../../common/util/Res"

export class ItemInfo {
    public ID: number// 物品id
    public Name: string// 物品名字
    public Type: number// 物品类型
    public Quality: number// 品质
}

export class ItemConfig {
    private static instance: ItemConfig
    public static Get(): ItemConfig {
        if (ItemConfig.instance == null) {
            ItemConfig.instance = new ItemConfig()
            ItemConfig.instance.init()
        }
        return ItemConfig.instance
    }

    public ItemSlc: Array<ItemInfo>
    public ItemMap: Map<number, ItemInfo>

    private init(): void {
        this.ItemSlc = new Array<ItemInfo>()
        this.ItemMap = new Map<number, ItemInfo>()

        let jsonData = Res.get<cc.JsonAsset>("json/Item", cc.JsonAsset)
        this.ItemSlc = jsonData.json['Item']
        this.ItemSlc.forEach(item => {
            this.ItemMap.set(item.ID, item)
        })

    }
}
