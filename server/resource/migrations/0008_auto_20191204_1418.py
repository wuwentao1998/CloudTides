# Generated by Django 2.2.7 on 2019-12-04 06:18

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('resource', '0007_auto_20191204_1000'),
    ]

    operations = [
        migrations.AlterField(
            model_name='resource',
            name='user',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL),
        ),
    ]
