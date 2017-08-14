%{!?BUILD_NUMBER: %define BUILD_NUMBER 1}
Name:           tcpstream
Version:        1.0.0
Release:        081401
Group:          Applications/Internet
Summary:        TCP stream proxy
License:        MIT License
URL:            https://www.livejournal.com
Source0:        tcpstream.tar.gz
BuildRoot:      %{_tmppath}/%{name}-%{version}-root-%(%{__id_u} -n)
Requires:       daemonize
Requires(pre):  shadow-utils

# pull in golang libraries by explicit import path, inside the meta golang()
# [...]

%description
# include your full description of the application here.

%prep
%setup -q -n %{name}

# many golang binaries are "vendoring" (bundling) sources, so remove them. Those dependencies need to be packaged independently.
rm -rf vendor

%build
export GOPATH=$(pwd)
export
go get -d
go build -v -a -ldflags "-B 0x$(head -c20 /dev/urandom|od -An -tx1|tr -d ' \n')" -tags 'netgo'

%install
rm -rf %{buildroot}
install -d %{buildroot}
install -d %{buildroot}%{_bindir}
install -d %{buildroot}%{_sysconfdir}
install -d  %{buildroot}%{_sysconfdir}/init.d
install -d  %{buildroot}/var/log/tcpstream
install -p -m 0755 ./tcpstream %{buildroot}%{_bindir}/tcpstream
install -p -m 0755 ./config.json.example %{buildroot}%{_sysconfdir}/tcpstream.json.example
install -p -m 0755 ./init.d %{buildroot}%{_sysconfdir}/init.d/tcpstream.example

%files
%defattr(-,root,root,-)
%attr(0755,root,root) %{_bindir}/tcpstream
%attr(0755,root,root) %{_sysconfdir}/tcpstream.json.example
%attr(0755,root,root) %{_sysconfdir}/init.d/tcpstream.example

%pre

%changelog

* Mon Aug 14 2017 Mikhail Kirillov <mikkirillov@yandex.ru> - 1.0.0
 - first package version

