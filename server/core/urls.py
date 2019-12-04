"""server URL Configuration
The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include, re_path
from django.views import generic
from django.conf.urls import url
from rest_framework_swagger.views import get_swagger_view
from django.conf.urls.static import static
from django.conf import settings

schema_view = get_swagger_view(title='TIDES API DOCS')

urlpatterns = [
    path('admin/', admin.site.urls),
    url(r'^api/docs/$', schema_view),
    path('api/users/', include('users.urls')),
    path('api/resource/', include('resource.urls')),
    path('api/policy/', include('policy.urls')),
    path('api/template/', include('template.urls')),
    path('api/usage/', include('usage.urls'))
]