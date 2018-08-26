# Cooperation with fzf.

function trash() {
  typeset -r trash="${HOME}/.Trash"
  local fzf_option="--preview-window='right:hidden' --bind='ctrl-v:toggle-preview'"

  ! type fzf > /dev/null 2>&1 && [[ -e ${GOPATH}/bin/trash ]] && ${GOPATH}/bin/trash $@

  case $1 in
    'move')
      [[ -z $2 ]] && set 'move' $(command ls -A ./ | sed "/^${trash##*/}$/"d \
        | eval "fzf --header='move files in the current directory to the trash' \
        --preview=\"file {} | sed 's/^.*: //'; du -hs {} | cut -f1; less {}\" ${fzf_option}") \
        > /dev/null && [[ -z $2 ]] && return
    ;;
    'restore')
      [[ -z $2 ]] && set 'restore' $(command ls -rA ${trash} \
        | eval "fzf --header='move files in the trash to the current directory' \
        --preview=\"file ${trash}/{} | sed 's/^.*: //'; du -hs ${trash}/{} | cut -f1; echo '\n'; less ${trash}/{}\" ${fzf_option}") \
        > /dev/null && [[ -z $2 ]] && return
    ;;
    'delete')
      [[ -z $2 ]] && set 'delete' $(command ls -rA ${trash} \
        | eval "fzf --header='delete files in the trash' \
        --preview=\"file ${trash}/{} | sed 's/^.*: //'; du -hs ${trash}/{} | cut -f1; echo '\n'; less ${trash}/{}\" ${fzf_option}") \
        > /dev/null && [[ -z $2 ]] && return
    ;;
    *)
    ;;
  esac

  [[ -e ${GOPATH}/bin/trash ]] && ${GOPATH}/bin/trash $@
}