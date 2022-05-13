# Terraform C2 Provider

Адаптация [Terraform AWS Provider](https://github.com/hashicorp/terraform-provider-aws) от HashiCorp под Облако КРОК (C2).

## Требования

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.26+
-	[Go](https://golang.org/doc/install) 1.17+

## Общая информация

**Провайдер** - это плагин для Terraform, который реализует возможность управления ресурсами некого сервиса 
(например, облака или БД) через его API. Для каждого ресурса в провайдер добавляется схема и CRUD операции.
Информация об API предоставляется в виде отдельной библиотеки. 

**Terraform С2 Provider** реализуется на базе **Terraform AWS Provider**. Также создан [форк](https://github.com/C2Devel/aws-sdk-go) 
библиотеки c AWS API.

Для публикации провайдера в [terraform registry](https://registry.terraform.io/) под новым именем (ранее - *aws*) форк
был переименован в *terraform-provider-croccloud*.

[ссылка на provider в registry](!)

## Начало работы

Для работы с провайдером требуется установка [Go](https://golang.org/doc/install) (см. [требования](#требования)).

Клонирование репозитория и сборка провайдера:
```
$ git clone git@github.com:C2Devel/terraform-provider-croccloud.git && cd terraform-provider-croccloud
...
$ make build
```

После сборки артефакт `terraform-provider-croccloud` будет доступен в директории `$GOPATH/bin`.

Получение значения `$GOPATH`: `go env GOPATH`

**Опционально.** Установка дополнительных библиотек (линтеры, форматтеры и т.д.): `make tools`

Установленные библиотеки также будут находиться в `$GOPATH/bin`.

## Тесты

В проекте есть два типа тестов: **unit** и **acceptance**. Они лежат рядом с функционалом (файлы: *_test.go). 

Тесты написаны с помощью [go testing](https://go.dev/doc/code#Testing), для приемочных дополнительно используется 
пакет [acctest](https://github.com/C2Devel/terraform-provider-croccloud/tree/main/internal/acctest).

### Unit

Запуск unit тестов: `make test`

### Acceptance

Для запуска приемочных тестов требуется установка [Terraform](https://golang.org/doc/install) (см. [требования](#требования)).

**Важно!** Тесты используют реальные облачные ресурсы. **Требуется доработка тестов для запуска на C2.**

По запуску и написанию приемочных тестов есть [документация](https://github.com/C2Devel/terraform-provider-croccloud/blob/main/docs/contributing/running-and-writing-acceptance-tests.md).

Команды:
- запуск всех тестов: `make testacc`
- запуск конкретного теста: `make testacc TESTS=TestAccEC2EBSVolume_basic PKG=ec2`

## Документация

Информация о провайдере расположена в директориях:
- `docs/` - инструкции для разработки, roadmap;
- `website/` - документация к провайдеру, которая публикуется в terraform registry 
([инструкция](https://www.terraform.io/registry/providers/docs) по работе с документацей от Terraform).

Структура директории website:
```
website/
|-- docs/ 
|    |-- d/                          # набор описаний для terraform data sources
|    |-- guides/
|    |-- r/                          # набор описаний для terraform resources
|    |    |-- <resource>.html.markdown
|    |    |-- ...
|    |
|    |-- index.html.markdown         # стартовая страница
|-- allowed-subcategories.txt        # разделы документации
```

## Выпуск релиза

Релиз провайдера формируется в соответствии с требованиями к публикации на terraform registry 
([инструкция](https://www.terraform.io/registry/providers/publishing#creating-a-github-release) по созданию github релиза от Terraform)

Подготовка релизных артефактов (сборка под разные архитектуры и ОС, подпись, архивация и т.д.) описана в `.goreleaser.yml`

### Настройка окружения

1. Установка [goreleaser](https://goreleaser.com/install/)

goreleaser будет установлен в `$GOPATH/bin`. Проверка установки:
```
$ export GOBIN=$(go env GOPATH)/bin
$ $GOBIN/goreleaser -v
...
```
2. Создание **GPG** ключа ([инструкция от github](https://docs.github.com/en/authentication/managing-commit-signature-verification/generating-a-new-gpg-key))
   - **Опционально.** Привязка ключа к github аккаунту
3. Установка переменной `GPG_FINGERPRINT`
```
$ gpg --list-secret-keys --keyid-format=long
...
sec   rsa4096/<id> 2022-04-19 [SC]
      <fingerprint>
...
$ export GPG_FINGERPRINT=<fingerprint>
```
4. Создание **Personal Access Token** с разрешением **public_repo** ([инструкция от github](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token))
5. Установка переменной `GITHUB_TOKEN`
```
$ export GITHUB_TOKEN=<generated token by github>
```

### Версионирование

Версии провайдера должны соответствовать спецификации [Semantic Version](https://semver.org/) и начинаться с **v** 
(например, **v1.2.3** или **v1.2.3-pre**). Версия провайдера фиксируется только в виде тэга. 

**Важно!** Не допускается обновление уже выпущенных версий, т.к. могут возникнуть проблемы со скачиванием провайдера из terraform registry. 

### Релиз

1. Запук тестов (см. [тесты](#тесты))
2. **Опционально.** Отключение автопубликации релиза в github (флаг `release.draft` в `.goreleaser.yml`)
```
# .goreleaser.yml
...
release:
  draft: true
```
3. Подготовка репозитория: на релизной ветке не должно быть незакоммиченных изменений, untracked файлов
4. Установка релизного тега с версией (см. [версионирование](#версионирование)) и его публикация
```
$ git tag v1.2.3
$ git push <remote> v1.2.3 
```
5. Сборка и подпись релизных артефактов. Артефакты будут размещены в директории `dist/`
```
$ $GOBIN/goreleaser release --rm-dist
```
6. **Опционально.** Если на шаге 2 была отключена автопубликация: создание релиза на github и загрузка артефактов:
   - `dist/terraform-provider-croccloud_{VERSION}_{OS}_{ARCH}.zip`
     - Для всех архитектур и ОС.
   - `dist/terraform-provider-croccloud_{VERSION}_SHA256SUMS`
   - `dist/terraform-provider-croccloud_{VERSION}_SHA256SUMS.sig`
   - `terraform-provider-croccloud_{VERSION}_manifest.json`
     - Файл создается вручную: `cp terraform-registry-manifest.json terraform-provider-croccloud_{VERSION}_manifest.json`

## Публикация провайдера в публичном registry

todo

https://www.terraform.io/registry/providers/publishing
gpg

Сначала надо выпустить версию

Удаление версий

## Публикация провайдера в private registry

todo

## Использование провайдера

## TODO

github actions (кажется, не прогоняют тесты, только линтеры)
линтеры настроить на пуш в релизную ветку (main ?)

версия aws-sdk-go
Починить приемочные тесты: выбор тестов, настройка тестового провайдера на C2 API + sweep

нужно ли для чего-то generate

возможно где-то потребуется установка GOBIN или добавление в PATH - линтеры?

в goreleaser 

выбор архитектур в goreleaser

**Важно!** Запуск тестов не включен в релизный процесс (github actions)

Проверить Using the provider


