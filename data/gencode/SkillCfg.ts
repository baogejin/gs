import Res from "../../common/util/Res"

export class SkillInfo {
    public ID: number// 技能id
    public Name: string// 技能名称
    public Level: number// 技能等级
    public BeforeTime: number// 前摇时间ms
    public AfterTime: number// 后摇时间ms
    public IsWeaponSkill: boolean// 是否为武器技能
    public CoolDown: number// 冷却时间ms
    public TargetType: number// 目标选择类型
    public Attack: number// 伤害量
    public Heal: number// 治疗量
}

export class SkillConfig {
    private static instance: SkillConfig
    public static Get(): SkillConfig {
        if (SkillConfig.instance == null) {
            SkillConfig.instance = new SkillConfig()
            SkillConfig.instance.init()
        }
        return SkillConfig.instance
    }

    public SkillSlc: Array<SkillInfo>
    public SkillMap: Map<number, SkillInfo>

    private init(): void {
        this.SkillSlc = new Array<SkillInfo>()
        this.SkillMap = new Map<number, SkillInfo>()

        let jsonData = Res.get<cc.JsonAsset>("json/Skill", cc.JsonAsset)
        this.SkillSlc = jsonData.json['Skill']
        this.SkillSlc.forEach(skill => {
            this.SkillMap.set(skill.ID, skill)
        })

    }
}
