#!/usr/bin/env bash

# On Debian and Ubuntu, installs the packages in requirements_apt.txt.
# Lines starting with '#' in requirements_apt.txt are ignored.
#
# On other platforms, errors out. Consider submitting a patch.

REPO_ROOT="$(git rev-parse --show-toplevel 2>/dev/null)" || {
  echo "Run this from the git repo of darktile.";
  exit 1;
}

# TODO: change this so the script looks for requirements_apt.txt in its
# own directory.
DEFAULT_REQUIREMENTS_APT_TXT_PATH="${REPO_ROOT}/scripts/requirements_apt.txt"

install_apt_requirements () {
  requirements_apt_path="$1"
  ok=1

  echo "# Installing build requirements on Debian";
  echo
  echo "# /etc/debian_version"
  cat /etc/debian_version
  echo
  echo "# /etc/lsb-release"
  cat /etc/lsb-release
  echo
  [[ -z "${requirements_apt_path}" ]] && {
    echo "InternalError; path to requirements_apt.txt not supplied."
    ok=0
  } || {
    [[ -f "${requirements_apt_path}" ]] && {
      echo "# requirements_apt.txt at ${requirements_apt_path}"
      cat "${requirements_apt_path}"
      echo
    } || {
      echo "RuntimeError; path to requirement_apt.txt does not exist";
      ok=0
    }
  }

  [[ "${ok}" == "1" ]] && {
    _install_apt_requirements "${requirements_apt_path}"
  } || {
    false  # return failure status from this function.
  }
}

# Core logic of install_apt_requirements; doesn't check arguments;
# Should be called only from install_apt_requirements().
#
# Does NOT handle spaces within package names aka Bash quoting hell.
_install_apt_requirements () {
  requirements_apt_path="$1"
  sudo apt install -y $(cat "${requirements_apt_path}")
}

install_build_requirements () {
  [[ -f /etc/debian_version ]] && {
    install_apt_requirements "${DEFAULT_REQUIREMENTS_APT_TXT_PATH}"
  } || {
    echo "Not running on debian. Not installing build requirements.";
    echo "Consider contributing a patch.";
  }
}

