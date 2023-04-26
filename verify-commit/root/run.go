package root

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.tap4fun.com/k2/githooks/aid"
	git "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	errGetCachedFiles = "[%s] 取准备提交的文件列表失败：%s"
	errInvalidBranch  = "[%s] 非法的分支名%s，您可能处于detached状态"
	errMustHave       = "[%s] 当前分支%s必须提交以下文件：\n%s"
	errRejected       = "[%s] 当前分支%s不允许提交以下文件：\n%s"
)

func Run(logger *cobra.Command, vpr *viper.Viper, hookName string, verbose bool) error {

	// extract the configure
	cfg, err := extractConfig(vpr.Get(hookName))
	if err != nil {
		return err
	}

	if verbose {
		logger.Printf("configure: %+v\n", cfg)
	}

	// open repository
	if err := os.Chdir(vpr.GetString("working-tree")); err != nil {
		return err
	}
	repo, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	// get current branch
	head, err := repo.Head()
	if err != nil {
		return err
	}
	if !head.Name().IsBranch() {
		return fmt.Errorf(errInvalidBranch, hookName, head.Name().String())
	}
	branch := head.Name().Short()
	if verbose {
		logger.Println("current branch:", branch)
	}

	staged, untracked, err := aid.GitStatus()
	if err != nil {
		return fmt.Errorf(errGetCachedFiles, hookName, err)
	}
	if len(staged) <= 0 {
		return nil
	}

	var approved, rejected []string
	includeOnlyCfg := cfg.IncludeOnly[branch]
	// 此分支的配置不存在则允许提交
	if includeOnlyCfg == nil {
		approved = staged
	} else {
		approved, rejected = match(includeOnlyCfg, staged)
		if len(rejected) > 0 {
			return fmt.Errorf(errRejected, hookName, branch, strings.Join(rejected, "\n"))
		}
	}

	// 是否要做Unity生产的meta文件的校验
	if !cfg.UnityMeta || len(untracked) <= 0 {
		return nil
	}

	// 新meta文件的基文件必须是
	//   1、允许提交的暂存区文件
	//   2、已经提交的文件或目录（TODO: 没找到方法判断是否已提交）
	approvedMap := aid.SliceToSet(approved, true, true)
	var must []string
	for _, metaFile := range untracked {
		lcMetaFile := strings.ToLower(strings.TrimSpace(metaFile))
		ext := filepath.Ext(lcMetaFile)
		if ext != ".meta" {
			continue
		}
		parentFile := strings.TrimSuffix(lcMetaFile, ".meta")

		// 基文件是允许提交的暂存区文件
		if ok := approvedMap[parentFile]; ok {
			must = append(must, metaFile)
			continue
		}

		// TODO 不知道如何根据文件名判断是否已提交
		//obj, err := repo.Object(plumbing.AnyObject, plumbing.NewHash(parentFile))
		//if err != nil {
		//	continue
		//}
		//
		//// 基文件是已提交的文件或目录
		//if obj.Type() == plumbing.BlobObject || obj.Type() == plumbing.TreeObject {
		//	must = append(must, metaFile)
		//}
	}

	if len(must) > 0 {
		return fmt.Errorf(errMustHave, hookName, branch, strings.Join(must, "\n"))
	}

	return nil
}

func match(cfg map[string]map[string]bool, staged []string) (approved, rejected []string) {
	for _, file := range staged {
		pass := false
		lowerCaseFile := strings.ToLower(strings.Trim(file, " \""))
		for dir, allowedExt := range cfg {
			if dir != "*" {
				dir = strings.ToLower(strings.TrimSpace(dir)) + "/*"
			}
			if ok, _ := filepath.Match(dir, lowerCaseFile); ok {
				// 此文件夹下的任意类型文件均可提交
				if _, ok := allowedExt["*"]; ok {
					pass = true
					break
				}

				// 扩展名可能为空，空扩展名也需要在配置里写明
				if allowedExt[filepath.Ext(lowerCaseFile)] {
					pass = true
					break
				}
			}
		}

		if pass {
			approved = append(approved, file)
		} else {
			rejected = append(rejected, file)
		}
	}
	return
}
