# SRS 计划

以 [SRS 评审稿](../requirements/SRS.md) 和 [用户故事汇总](../user-stories/user-stories-summary.md) 为准。各负责人先核对分配给自己的 GitHub Issue，课前完成用例细化、用例图校对和 SRS 草稿 PR。

| 负责人 | GitHub | 功能用例 | 非功能验收 |
| --- | --- | --- | --- |
| 赵睿妍 | ruiyan15 | UC-01 至 UC-05 | NFC-04、NFC-05 |
| 郭芷烟 | CasttiFlorenz | UC-07、UC-09 至 UC-11 | 汇总 SRS 与用例图 |
| 赵晨瑞 | zcr324 | UC-12 至 UC-15 | NFC-01、NFC-02 |
| 张子原 | Shanyou-Fusu | UC-08、UC-16 至 UC-19 | NFC-03 |
| 于恩熙 | StellarX12 | UC-06、UC-20 至 UC-23 | NFC-06 |

1. **细化用例。** 在 SRS 对应用例小节补齐角色、触发条件、前置条件、主流程、关键异常、结果和验收标准。保持 UC、FR 编号，关联故事 Issue。跨模块规则先在 Issue 讨论，例如互评资格、匿名权限和补派后的迟交。
2. **细化非功能验收。** 按各自 Issue 写出测试账号或数据、操作步骤、预期结果和失败判据。性能测试记录数据量和环境，匿名与权限测试覆盖授权与无关账号。SRS 中 NFC-07 的追溯引用待评审确认，当前第 5 节已有定义的是 NFC-01 至 NFC-06。
3. **提交与评审。** 从最新 main 开分支，提交本人负责的 SRS 小节。PR 用 Refs 关联用例 Issue，正文列出改动节号、验收场景和待决规则。组员检查编号、权限和主异常流程后批准合并。用例图由郭芷烟汇总，其他成员核对自己的角色与用例关系。
4. **课堂与基线。** 第 3 课前准备完整用例图和 SRS 草稿 PR，汇报关键场景及非功能通过条件。课堂评审通过后归档追溯矩阵 v0 与变更记录，发布 SRS v1.0。后续变更走 Issue 申请、评审批准和关联 PR。

## Issue 索引

| 编号 | GitHub Issue | 负责人 |
| --- | --- | --- |
| R01 | [#14](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/14) | ruiyan15 |
| R02 | [#15](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/15) | ruiyan15 |
| R03-1 | [#16](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/16) | CasttiFlorenz |
| R03-2 | [#17](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/17) | CasttiFlorenz |
| R04-1 | [#18](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/18) | CasttiFlorenz |
| R04-2 | [#19](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/19) | CasttiFlorenz |
| R05 | [#11](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/11) | zcr324 |
| R06 | [#12](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/12) | zcr324 |
| R07 | [#20](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/20) | Shanyou-Fusu |
| R08 | [#21](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/21) | Shanyou-Fusu |
| R09 | [#22](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/22) | StellarX12 |
| R10 | [#23](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/23) | StellarX12 |
| UC-01 | [#24](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/24) | ruiyan15 |
| UC-02 | [#25](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/25) | ruiyan15 |
| UC-03 | [#26](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/26) | ruiyan15 |
| UC-04 | [#27](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/27) | ruiyan15 |
| UC-05 | [#28](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/28) | ruiyan15 |
| UC-06 | [#29](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/29) | StellarX12 |
| UC-07 | [#30](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/30) | CasttiFlorenz |
| UC-08 | [#31](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/31) | Shanyou-Fusu |
| UC-09 | [#32](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/32) | CasttiFlorenz |
| UC-10 | [#33](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/33) | CasttiFlorenz |
| UC-11 | [#34](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/34) | CasttiFlorenz |
| UC-12 | [#35](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/35) | zcr324 |
| UC-13 | [#36](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/36) | zcr324 |
| UC-14 | [#37](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/37) | zcr324 |
| UC-15 | [#38](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/38) | zcr324 |
| UC-16 | [#39](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/39) | Shanyou-Fusu |
| UC-17 | [#40](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/40) | Shanyou-Fusu |
| UC-18 | [#41](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/41) | Shanyou-Fusu |
| UC-19 | [#42](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/42) | Shanyou-Fusu |
| UC-20 | [#43](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/43) | StellarX12 |
| UC-21 | [#44](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/44) | StellarX12 |
| UC-22 | [#45](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/45) | StellarX12 |
| UC-23 | [#46](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/46) | StellarX12 |
| NFC-01 | [#47](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/47) | zcr324 |
| NFC-02 | [#48](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/48) | zcr324 |
| NFC-03 | [#49](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/49) | Shanyou-Fusu |
| NFC-04 | [#50](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/50) | ruiyan15 |
| NFC-05 | [#51](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/51) | ruiyan15 |
| NFC-06 | [#52](https://github.com/CasttiFlorenz/software-engineering-course-design/issues/52) | StellarX12 |
