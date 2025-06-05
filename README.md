# Лабораторная работа по распределённым системам (ITMO / Яндекс.Практикум)

Репозиторий создан для изучения архитектуры распределённых систем с использованием облачной платформы Yandex Cloud. В проекте разворачивается инфраструктура для системы с балансировкой нагрузки между мастер- и слейв-серверами PostgreSQL.

## 🛠 Технологии
- **Инфраструктура**: Terraform + Yandex Cloud
- **Конфигурация**: Ansible
- **СУБД**: PostgreSQL (мастер-слейв репликация)

## 🚀 Быстрый старт

### 1. Предварительные требования
- Аккаунт в [Yandex Cloud](https://cloud.yandex.ru/)
- Установленный [Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli) (1.11.3)
- Установленный [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/index.html) (2.15.13)
- Доступ к YC CLI (`yc`)

### 2. Настройка переменных
перейдите в папку infrs
```bash

cd infra
```

Скопируйте шаблон переменных и укажите свои значения:
```bash

cp terraform_exp.tfvars terraform.tfvars
nano terraform.tfvars  # отредактируйте файл
```

Или сразу создайте папку secrets, скопируйте и отредактируйте файл там
```bash
mkdir secrets
cp terraform_exp.tfvars secrets/terraform.tfvars
nano secrets/terraform.tfvars
```

### 3. Развёртывание инфраструктуры
```bash

terraform init
terraform apply -auto-approve # или terraform apply --var-file="secrets/terraform.tfvars" -auto-approve
```
После выполнения:
- В `infra/secrets/vm_ips.json` появятся IP-адреса ВМ
- Автоматически генерируется инвентарный файл для Ansible

### 4. Настройка файла inventory.ini
Воспользуйтесь файлом `example_inventory.ini`. В нём есть всё необходимое для запуска пайплайна ansible.
Замените адреса виртуальных машин на ваши собственные, если вы пользуетесь общей внутренней сетью (например в яндекс облаке)
используйте в переменной `db_master_internal_host` внутренний адрес, в противном случае можно использовать внешний, такой же как в переменной `db_master`

### 5. Настройка серверов через Ansible
```bash

cd ansible/
ansible-playbook -i inventory.yml playbook.yml

📂 Структура репозитория
├── ansible
│   ├── ansible.cfg
│   ├── example_inventory.ini // пример файла inventory
│   ├── group_vars
│   │   └── ... // Переменные для ролей
│   ├── inventory.ini // отсутствует в репозитории, воспользуйтесь примером
│   ├── playbook.yml
│   ├── roles
│   │   └── ... // Роли по шагам
├── backend
│   └── ... // Реализация backend
├── infra
│   └── ... // Реализация инфраструктуры на terraform
├── LICENSE
└── README.md

```

🔐 Безопасность
- Все sensitive-данные (токены, пароли) исключены из репозитория через .gitignore
- Для доступа к ВМ используются только SSH-ключи

📝 Лицензия
MIT License. Подробнее в файле LICENSE.

> **Note**
> 
> Проект разработан в рамках учебного курса "Распределённые системы" в Университете ИТМО при поддержке Яндекс.Практикума.