package cmd

import (
	"context"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/izumin5210/gex/pkg/tool"
	"github.com/izumin5210/grapi/pkg/excmd"
	"github.com/izumin5210/grapi/pkg/gogen"
	"github.com/izumin5210/grapi/pkg/gogen/di"
	"github.com/izumin5210/grapi/pkg/gogen/util/fs"
)

func newGenerateCommands(ctx *gogen.Ctx) (cmds []*cobra.Command) {
	gCmd := &cobra.Command{
		Use:     "generate GENERATOR",
		Short:   "Generate a new code",
		Aliases: []string{"g", "gen"},
	}
	dCmd := &cobra.Command{
		Use:     "destroy GENERATOR",
		Short:   "Destroy an existing new code",
		Aliases: []string{"d"},
	}

	var (
		execs []string
		wg    sync.WaitGroup
	)

	wg.Add(2)
	cmdNames := make(map[string]struct{}, 100)

	go func() {
		defer wg.Done()
		execs = fs.ListExecutableWithPrefix(ctx.FS, "grapi-gen-")
	}()

	go func() {
		defer wg.Done()

		toolRepo, err := di.NewToolRepository(ctx)
		if err != nil {
			zap.L().Debug("failed to initialize tools repository", zap.Error(err))
			return
		}

		tools, err := toolRepo.List(context.TODO())
		if err != nil {
			zap.L().Debug("failed to retrieve tools", zap.Error(err))
			return
		}

		for _, t := range tools {
			if !strings.HasPrefix(t.Name(), "grapi-gen-") {
				continue
			}
			if _, ok := cmdNames[t.Name()]; ok {
				continue
			}
			cmdNames[t.Name()] = struct{}{}
			gCmd.AddCommand(newGenerateCommandByTool(toolRepo, t, "generate"))
			dCmd.AddCommand(newGenerateCommandByTool(toolRepo, t, "destroy"))
		}
	}()

	wg.Wait()

	for _, exec := range execs {
		if _, ok := cmdNames[exec]; ok {
			continue
		}
		cmdNames[exec] = struct{}{}
		gCmd.AddCommand(newGenerateCommandByExec(di.NewCommandExecutor(ctx), exec, "generate"))
		dCmd.AddCommand(newGenerateCommandByExec(di.NewCommandExecutor(ctx), exec, "destroy"))
	}

	cmds = append(cmds, gCmd, dCmd)

	return
}

func newGenerateCommandByTool(repo tool.Repository, t tool.Tool, subCmd string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  strings.TrimPrefix(t.Name(), "grapi-gen-"),
		Args: cobra.ArbitraryArgs,
		RunE: func(_ *cobra.Command, args []string) error {
			return repo.Run(context.TODO(), t.Name(), append([]string{subCmd}, args...)...)
		},
	}
	cmd.SetUsageFunc(func(*cobra.Command) error {
		return repo.Run(context.TODO(), t.Name(), subCmd, "--help")
	})
	return cmd
}

func newGenerateCommandByExec(execer excmd.Executor, exec, subCmd string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  strings.TrimPrefix(exec, "grapi-gen-"),
		Args: cobra.ArbitraryArgs,
		RunE: func(_ *cobra.Command, args []string) error {
			_, err := execer.Exec(context.TODO(), exec, excmd.WithArgs(append([]string{subCmd}, args...)...), excmd.WithIOConnected())
			return err
		},
	}
	cmd.SetUsageFunc(func(*cobra.Command) error {
		_, err := execer.Exec(context.TODO(), exec, excmd.WithArgs(subCmd, "--help"), excmd.WithIOConnected())
		return err
	})
	return cmd
}