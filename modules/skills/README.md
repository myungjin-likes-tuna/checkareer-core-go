# skills

스킬을 관리하는 모듈

## constraints

```
CREATE CONSTRAINT unique_skill FOR (skill:Skill) REQUIRE skill.name IS UNIQUE
```
