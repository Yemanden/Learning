**Adapter**

Реализация паттерна Adapter.
Состоит из 4 основных частей:
- Экран
- Читатель
- Бумажный документ
- Электронный документ
И 1 вспомогательной части:
- Конвертер

Логика реализации: читатель может легко прочитать бумажный документ, читать электронный документ он не умеет. Для того, чтобы прочитать эл. документ ему понадобится некий адаптер,
который сможет конвертировать документ в пригодный для чтения вид и отобразит его.
В роли адаптера выступает экран.